#ifndef AIU_NET_H
#define AIU_NET_H

#include <stdint.h>

#define MAX_NET_MESSAGES 64
#define MAX_LEN_OVERHEAD 12
#define MAX_LEN_PAYLOAD  8
#define MAX_LEN_FRAME	(MAX_LEN_PAYLOAD + MAX_LEN_OVERHEAD)
#define MAX_COM_CHANNEL  2

#define NETWORK_DEFAULT_UID	555
#define NETWORK_UID 		555
#define OWN_NET_ADDR 		0xFF20
#define OWN_BIND_ID		0

#define THIS_NET_ADDR 		0x0000
#define MAX_OWN_NET_ADDR 	0x7FFF
#define PINCODE_NET_CONFIG 	67890
#define PINCODE_EXTNET_CONFIG	123

#define MIN_UI_NET_ADDR 	0x8000
#define MAX_UI_NET_ADDR 	0xFEFF

#define BROADCAST_LAN_ADDR 	0xFF00 // All RF-local units
#define BROADCAST_RF 		BROADCAST_LAN_ADDR
#define CHANNEL_RF 			BROADCAST_RF

#define BROADCAST_NET_ADDR 	0xFF01 // All RF-local and gates external units
#define BROADCAST_RF_AND_EXTERNAL BROADCAST_NET_ADDR
//RF and any external gates
//...

#define BROADCAST_WAN_ADDR 	0xFF10 // All gates external units only, NOT RF-local units
#define BROADCAST_EXTERNAL 	BROADCAST_WAN_ADDR
#define CHANNEL_EXTERNAL 	BROADCAST_EXTERNAL

//*** Addresses gates devises
#define BLUETOOTH_NET_ADDR 0xFF20
#define WIFI_NET_ADDR 0xFF21
#define ETHERNET_NET_ADDR 0xFF22
#define USB_NET_ADDR 0xFF23

#define NOT_FOUND 0xFF

#define	MSG_INTERNAL_SYNCRO	51

#define PACKET_ID_SIGNAL 14
#define PACKET_ID_SHORTLIVED_PIN 31
#define PACKET_ID_TERMINAL_CMD 32


typedef enum{
	CHANNEL_STATE_OFF   	= 0x00,
	CHANNEL_STATE_READY  	= 0x01,
	CHANNEL_STATE_BUSY   	= 0x02,
    CHANNEL_STATE_OPERATES 	= 0x02,
	//CHANNEL_STATE_LOADSCRIPT= 0x03,
	CHANNEL_STATE_SLEEP   	= 0x04,
} comChannelState_t;

typedef enum{
	NETMSG_ATTR_FREE	= 0,
    NETMSG_ATTR_BUSY,
    NETMSG_ATTR_AIUFORMAT,
    NETMSG_ATTR_DATAFORMAT
} netmsgAttr_t;

typedef enum{
	CHANNEL_ATTR_SENDED   	= 0x00,
        CHANNEL_ATTR_NEEDSEND	= 0x01,
} channelAttr_Sends_t;

typedef enum{
	CHANNEL_ATTR_NOTCONFIRM	    = 0x00,
        CHANNEL_ATTR_WAITCONFIRM    = 0x01,
} channelAttr_Confirm_t;

// Up to 32 channels. Set bits... 
typedef enum{
	CHANNEL_AIU868  = ((uint16_t)0x0001),
	CHANNEL_AIU2400	= ((uint16_t)0x0002),
    CHANNEL_BT	= ((uint16_t)0x0004),
	CHANNEL_BLE	= ((uint16_t)0x0008),
    CHANNEL_WIFI	= ((uint16_t)0x0010),
    CHANNEL_3G	= ((uint16_t)0x0020),
	CHANNEL_UART0	= ((uint16_t)0x0040),
	CHANNEL_UART1	= ((uint16_t)0x0080),
	CHANNEL_UART2	= ((uint16_t)0x0100),
	CHANNEL_SPI0	= ((uint16_t)0x0200),
	CHANNEL_SPI1	= ((uint16_t)0x0400),
	CHANNEL_SPI2	= ((uint16_t)0x0800),
	CHANNEL_I2C0	= ((uint16_t)0x1000),
	CHANNEL_I2C1	= ((uint16_t)0x2000),
	CHANNEL_I2C2	= ((uint16_t)0x4000),

    CHANNEL_ALL		= ((uint16_t)0xFFFF),
} nameChannel_t;

typedef uint8_t	    payload_t[MAX_LEN_PAYLOAD];
typedef struct{
	uint8_t	    length;
	uint8_t	    id;
	uint16_t    destAddr;
	uint32_t    netUId;
	uint16_t    sourceAddr;
	uint16_t    misc;
	payload_t   payload;
	int8_t	    metricsRSSI;
	int8_t	    metricsLQI;

	uint8_t	    age;
        uint8_t	    attr_Status;
	uint16_t    attr_Send; // each bit - flag of appropriate channel need sends.
        uint16_t    attr_Confirm; // each bit - flag of appropriate channel need confirmation.
} netMessage_t;


void aiu_initNetMsgs(void);
static uint32_t getNewIndexNetMsg(void);
void aiu_addNetMsg(uint16_t channels, uint8_t packetId, uint16_t destAddr, uint8_t *buffer, uint8_t lenBuffer, uint16_t miscData);
void aiu_addNetMsg_fromBuffer  (uint16_t channels, uint8_t *buffer, uint8_t lenBuffer);
void aiu_addDataMsg_fromBuffer (uint16_t channels, uint8_t *buffer, uint8_t lenBuffer);
uint8_t aiu_getReadyNetMsg(uint16_t channels);
uint8_t aiu_getReadyDataMsg(uint16_t channels);
void aiu_eraseNetMsg(uint8_t idx);
void aiu_setChannelNetMsg(uint16_t channels, uint8_t idx);
void aiu_resetChannelNetMsg(uint16_t channels, uint8_t idx);
void aiu_agingNetMsg(void);

char * utoa16Simple(uint32_t value, char *buffer);


#endif

