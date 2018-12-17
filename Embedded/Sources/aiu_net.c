
#include "aiu_net.h"
#include "nrf_nvic.h"
#include <stdint.h>
#include <string.h>

netMessage_t listNetMsg[MAX_NET_MESSAGES];
static uint32_t _cnt = 0;
static volatile uint32_t _idx = 0;

void aiu_initNetMsgs() {

    memset(listNetMsg, 0, sizeof(listNetMsg));
}

// Return index of message in list.
static uint32_t getNewIndexNetMsg() {

    int index = 0;
    for (uint32_t i = 0; i < MAX_NET_MESSAGES; i++) {
        __sd_nvic_irq_disable();
        if (listNetMsg[i].attr_Status == NETMSG_ATTR_FREE) {
            listNetMsg[i].attr_Status = NETMSG_ATTR_BUSY;
            index = i;
            __sd_nvic_irq_enable();
            memset(&listNetMsg[i], 0, MAX_LEN_FRAME);
            break;
        } else {
            __sd_nvic_irq_enable();
        }
    }
	//aiu_agingNetMsg();
    return index;
}

// Len=1byte, ID=1byte, Net_A=4byte, Dest_A=2byte, Src_A=2byte, Misc=2byte. A buffer - data of payload.
void aiu_addNetMsg(uint16_t channels, uint8_t packetId, uint16_t destAddr, uint8_t *buffer, uint8_t lenBuffer, uint16_t miscData) {
    if (lenBuffer > MAX_LEN_FRAME)
        return;
    uint32_t index = getNewIndexNetMsg();

    listNetMsg[index].attr_Send = channels;
    listNetMsg[index].attr_Confirm = 0x00;
    listNetMsg[index].length = MAX_LEN_FRAME;
    listNetMsg[index].id = packetId;
    listNetMsg[index].destAddr = destAddr;
    listNetMsg[index].sourceAddr = OWN_NET_ADDR;
    listNetMsg[index].misc = miscData;
    listNetMsg[index].netUId = NETWORK_UID;
    //    if (listNetMsg[index].misc & FLAG_BINDUNIT) {
    //        listNetMsg[index].netUId = OWN_BIND_ID;
    //    } else {
    //        listNetMsg[index].netUId = NETWORK_UID;
    //    }

    memcpy(&listNetMsg[index].payload[0], buffer, lenBuffer);
    listNetMsg[index].age = MAX_NET_MESSAGES - 1;
    listNetMsg[index].attr_Status = NETMSG_ATTR_AIUFORMAT;
    aiu_agingNetMsg();
}

// A buffer - whole frame.
void aiu_addNetMsg_fromBuffer(uint16_t channels, uint8_t *buffer, uint8_t lenBuffer) {

    if ((lenBuffer) > MAX_LEN_FRAME)
        return;
    _cnt++;
    uint32_t index = getNewIndexNetMsg();
    _idx = lenBuffer;
    listNetMsg[index].attr_Send = channels;
    listNetMsg[index].attr_Confirm = 0x00;
    memcpy(&listNetMsg[index], buffer, lenBuffer);
    listNetMsg[index].age = MAX_NET_MESSAGES - 1;
    listNetMsg[index].attr_Status = NETMSG_ATTR_AIUFORMAT; // Maybe DATAFORMAT

    //!!!
    //    if (listNetMsg[index].payload[0] == 0x33){
    //	listNetMsg[index].age = MAX_NET_MESSAGES;
    //    }
}


// A buffer - whole frame.
void aiu_addDataMsg_fromBuffer(uint16_t channels, uint8_t *buffer, uint8_t lenBuffer)
{
    if ((lenBuffer) > MAX_LEN_FRAME)
        return;
    _cnt++;
    uint32_t index = getNewIndexNetMsg();
    _idx = lenBuffer;
    listNetMsg[index].attr_Send = channels;
    listNetMsg[index].attr_Confirm = 0x00;
    memcpy(&listNetMsg[index], buffer, lenBuffer);
    listNetMsg[index].age = MAX_NET_MESSAGES - 1;
    listNetMsg[index].attr_Status = NETMSG_ATTR_DATAFORMAT;
}


#define NOT_FOUND 0xFF
// Return index
uint8_t aiu_getReadyNetMsg(uint16_t channels) {

    uint8_t minNet = MAX_NET_MESSAGES;
    uint8_t idxNet = NOT_FOUND;
    for (int32_t i = 0; i < MAX_NET_MESSAGES; i++) {
        if (listNetMsg[i].attr_Status == NETMSG_ATTR_AIUFORMAT) {
            if ((listNetMsg[i].attr_Send & channels) != 0) {
                if (listNetMsg[i].age < minNet) {
                    minNet = listNetMsg[i].age;
                    idxNet = i;
                }
            }
        }
    }
    return idxNet;
}


// Return index
uint8_t aiu_getReadyDataMsg(uint16_t channels)
{
    uint8_t minNet = MAX_NET_MESSAGES;
    uint8_t idxNet = NOT_FOUND;
    for (int32_t i = 0; i < MAX_NET_MESSAGES; i++) {
        if (listNetMsg[i].attr_Status == NETMSG_ATTR_DATAFORMAT) {
            if ((listNetMsg[i].attr_Send & channels) != 0) {
                if (listNetMsg[i].age < minNet) {
                    minNet = listNetMsg[i].age;
                    idxNet = i;
                }
            }
        }
    }
    return idxNet;
}


void aiu_eraseNetMsg(uint8_t idx) {

    listNetMsg[idx].attr_Status = NETMSG_ATTR_FREE;
    listNetMsg[idx].attr_Send = 0x00;
}

void aiu_setChannelNetMsg(uint16_t channels, uint8_t idx) {

    listNetMsg[idx].attr_Send |= channels;
}

void aiu_resetChannelNetMsg(uint16_t channels, uint8_t idx) {

    listNetMsg[idx].attr_Send &= ~channels;
}

void aiu_agingNetMsg() {

    for (int32_t i = 0; i < MAX_NET_MESSAGES; i++) {
        //__sd_nvic_irq_disable();
        if (listNetMsg[i].attr_Status != NETMSG_ATTR_FREE) {
            if (listNetMsg[i].age == 0) {
                if ((listNetMsg[i].attr_Send == 0) && (listNetMsg[i].attr_Confirm == 0)) {
                    // Msg sended to all channels and not wait confirmation
                    listNetMsg[i].attr_Status = NETMSG_ATTR_FREE;
                }
            } else {
                listNetMsg[i].age--;
            }
        }
        //__sd_nvic_irq_disable();
    }
}


char * utoa16Simple(uint32_t value, char *buffer)
{
	static const char _line []="0123456789ABCDEF";
	buffer += 8;
	*--buffer = 0;
	do
	{
		*--buffer = _line[value % 16];
		value /= 16;
	}
	while (value);
	return buffer;
}


