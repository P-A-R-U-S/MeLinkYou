PROJECT_NAME     := meLinkYou
TARGETS          := nrf52832_xxaa
OUTPUT_DIRECTORY := _build

SDK_ROOT  := ./Sources/nRF5_SDK/nRF5_SDK_14.2.0_17b948a
#SDK_ROOT  := d:/nrf52/nRF5_SDK_14.2.0

JLINK := JLinkExe
JLINK_DIR := /Applications/SEGGER/JLink_V632g/
#JLINK_DIR := c:/JLink_V614a
PROJ_DIR  := .

$(OUTPUT_DIRECTORY)/nrf52832_xxaa.out: \
  LINKER_SCRIPT  := my_nrf52.ld

#VECTOR	= $(SDK_ROOT)/components/toolchain/gcc/gcc_startup_nrf52.S
VECTOR	= ./Sources/device/gcc_startup_nrf52.S


INC_FOLDERS += $(SDK_ROOT)/components \
	  $(SDK_ROOT)/components/drivers_nrf/comp \
	  $(SDK_ROOT)/components/drivers_nrf/twi_master \
	  $(SDK_ROOT)/components/ble/ble_services/ble_ancs_c \
	  $(SDK_ROOT)/components/ble/ble_services/ble_ias_c \
	  $(SDK_ROOT)/components/libraries/pwm \
	  $(SDK_ROOT)/components/softdevice/s132/headers/nrf52 \
	  $(SDK_ROOT)/components/libraries/usbd/class/cdc/acm \
	  $(SDK_ROOT)/components/libraries/usbd/class/hid/generic \
	  $(SDK_ROOT)/components/libraries/usbd/class/msc \
	  $(SDK_ROOT)/components/libraries/usbd/class/hid \
	  $(SDK_ROOT)/components/libraries/experimental_log \
	  $(SDK_ROOT)/components/ble/ble_services/ble_gls \
	  $(SDK_ROOT)/components/libraries/fstorage \
	  $(SDK_ROOT)/components/drivers_nrf/i2s \
	  $(SDK_ROOT)/components/libraries/mutex \
	  $(SDK_ROOT)/components/libraries/gpiote \
	  $(SDK_ROOT)/components/libraries/experimental_log/src \
	  $(SDK_ROOT)/components/drivers_nrf/gpiote \
	  $(SDK_ROOT)/components/libraries/fifo \
	  $(SDK_ROOT)/components/boards \
	  $(SDK_ROOT)/components/libraries/experimental_memobj \
	  $(SDK_ROOT)/components/drivers_nrf/common \
	  $(SDK_ROOT)/components/ble/ble_advertising \
	  $(SDK_ROOT)/components/ble/ble_services/ble_bas_c \
	  $(SDK_ROOT)/components/ble/ble_services/ble_hrs_c \
	  $(SDK_ROOT)/components/libraries/queue \
	  $(SDK_ROOT)/components/libraries/pwr_mgmt \
	  $(SDK_ROOT)/components/ble/ble_dtm \
	  $(SDK_ROOT)/components/toolchain/cmsis/include \
	  $(SDK_ROOT)/components/ble/ble_services/ble_rscs_c \
	  $(SDK_ROOT)/components/drivers_nrf/uart \
	  $(SDK_ROOT)/components/ble/common \
	  $(SDK_ROOT)/components/ble/ble_services/ble_lls \
	  $(SDK_ROOT)/components/drivers_nrf/wdt \
	  $(SDK_ROOT)/components/ble/ble_services/ble_bas \
	  $(SDK_ROOT)/components/libraries/experimental_section_vars \
	  $(SDK_ROOT)/components/softdevice/s132/headers \
	  $(SDK_ROOT)/components/ble/ble_services/ble_ans_c \
	  $(SDK_ROOT)/components/libraries/slip \
	  $(SDK_ROOT)/components/libraries/mem_manager \
	  $(SDK_ROOT)/external/segger_rtt \
	  $(SDK_ROOT)/components/libraries/usbd/class/cdc \
	  $(SDK_ROOT)/components/drivers_nrf/hal \
	  $(SDK_ROOT)/components/ble/ble_services/ble_nus_c \
	  $(SDK_ROOT)/components/drivers_nrf/rtc \
	  $(SDK_ROOT)/components/softdevice/common \
	  $(SDK_ROOT)/components/ble/ble_services/ble_ias \
	  $(SDK_ROOT)/components/libraries/usbd/class/hid/mouse \
	  $(SDK_ROOT)/components/libraries/ecc \
	  $(SDK_ROOT)/components/drivers_nrf/ppi \
	  $(SDK_ROOT)/components/ble/ble_services/ble_dfu \
	  $(SDK_ROOT)/external/fprintf \
	  $(SDK_ROOT)/components/drivers_nrf/twis_slave \
	  $(SDK_ROOT)/components/libraries/atomic \
	  $(SDK_ROOT)/components \
	  $(SDK_ROOT)/components/libraries/scheduler \
	  $(SDK_ROOT)/components/libraries/cli \
	  $(SDK_ROOT)/components/ble/ble_services/ble_lbs \
	  $(SDK_ROOT)/components/ble/ble_services/ble_hts \
	  $(SDK_ROOT)/components/drivers_nrf/delay \
	  $(SDK_ROOT)/components/libraries/crc16 \
	  $(SDK_ROOT)/components/drivers_nrf/timer \
	  $(SDK_ROOT)/components/libraries/util \
	  $(SDK_ROOT)/components/drivers_nrf/pwm \
	  $(SDK_ROOT)/components/libraries/csense_drv \
	  $(SDK_ROOT)/components/libraries/csense \
	  $(SDK_ROOT)/components/libraries/balloc \
	  $(SDK_ROOT)/components/libraries/low_power_pwm \
	  $(SDK_ROOT)/components/libraries/hardfault \
	  $(SDK_ROOT)/components/ble/ble_services/ble_cscs \
	  $(SDK_ROOT)/components/libraries/uart \
	  $(SDK_ROOT)/components/libraries/hci \
	  $(SDK_ROOT)/components/libraries/usbd/class/hid/kbd \
	  $(SDK_ROOT)/components/drivers_nrf/spi_slave \
	  $(SDK_ROOT)/components/drivers_nrf/lpcomp \
	  $(SDK_ROOT)/components/libraries/timer \
	  $(SDK_ROOT)/components/drivers_nrf/rng \
	  $(SDK_ROOT)/components/drivers_nrf/power \
	  $(SDK_ROOT)/components/libraries/usbd/config \
	  $(SDK_ROOT)/components/toolchain \
	  $(SDK_ROOT)/components/libraries/led_softblink \
	  $(SDK_ROOT)/components/drivers_nrf/qdec \
	  $(SDK_ROOT)/components/ble/ble_services/ble_cts_c \
	  $(SDK_ROOT)/components/drivers_nrf/spi_master \
	  $(SDK_ROOT)/components/ble/ble_services/ble_nus \
	  $(SDK_ROOT)/components/libraries/twi_mngr \
	  $(SDK_ROOT)/components/ble/ble_services/ble_hids \
	  $(SDK_ROOT)/components/libraries/strerror \
	  $(SDK_ROOT)/components/drivers_nrf/pdm \
	  $(SDK_ROOT)/components/libraries/crc32 \
	  $(SDK_ROOT)/components/libraries/usbd/class/audio \
	  $(SDK_ROOT)/components/ble/peer_manager \
	  $(SDK_ROOT)/components/drivers_nrf/swi \
	  $(SDK_ROOT)/components/ble/ble_services/ble_tps \
	  $(SDK_ROOT)/components/ble/ble_services/ble_dis \
	  $(SDK_ROOT)/components/device \
	  $(SDK_ROOT)/components/ble/nrf_ble_gatt \
	  $(SDK_ROOT)/components/ble/nrf_ble_qwr \
	  $(SDK_ROOT)/components/libraries/usbd \
	  $(SDK_ROOT)/components/drivers_nrf/saadc \
	  $(SDK_ROOT)/components/libraries/atomic_fifo \
	  $(SDK_ROOT)/components/ble/ble_services/ble_lbs_c \
	  $(SDK_ROOT)/components/ble/ble_racp \
	  $(SDK_ROOT)/components/toolchain/gcc \
	  $(SDK_ROOT)/components/libraries/fds \
	  $(SDK_ROOT)/components/libraries/twi \
	  $(SDK_ROOT)/components/drivers_nrf/clock \
	  $(SDK_ROOT)/components/ble/ble_services/ble_rscs \
	  $(SDK_ROOT)/components/drivers_nrf/usbd \
	  $(SDK_ROOT)/components/ble/ble_services/ble_hrs \
      ./Sources/config \
      ./Sources/device \
      ./Sources \
 
SRC_FILES += $(VECTOR) \
	  $(SDK_ROOT)/components/libraries/experimental_log/src/nrf_log_backend_rtt.c \
	  $(SDK_ROOT)/components/libraries/experimental_log/src/nrf_log_backend_serial.c \
	  $(SDK_ROOT)/components/libraries/experimental_log/src/nrf_log_default_backends.c \
	  $(SDK_ROOT)/components/libraries/experimental_log/src/nrf_log_frontend.c \
	  $(SDK_ROOT)/components/libraries/experimental_log/src/nrf_log_str_formatter.c \
	  $(SDK_ROOT)/components/libraries/util/app_error.c \
	  $(SDK_ROOT)/components/libraries/util/app_error_weak.c \
	  $(SDK_ROOT)/components/libraries/fifo/app_fifo.c \
	  $(SDK_ROOT)/components/libraries/scheduler/app_scheduler.c \
	  $(SDK_ROOT)/components/libraries/timer/app_timer.c \
	  $(SDK_ROOT)/components/libraries/uart/app_uart_fifo.c \
	  $(SDK_ROOT)/components/libraries/util/app_util_platform.c \
	  $(SDK_ROOT)/components/libraries/hardfault/hardfault_implementation.c \
	  $(SDK_ROOT)/components/libraries/util/nrf_assert.c \
	  $(SDK_ROOT)/components/libraries/atomic_fifo/nrf_atfifo.c \
	  $(SDK_ROOT)/components/libraries/balloc/nrf_balloc.c \
	  $(SDK_ROOT)/external/fprintf/nrf_fprintf.c \
	  $(SDK_ROOT)/external/fprintf/nrf_fprintf_format.c \
	  $(SDK_ROOT)/components/libraries/fstorage/nrf_fstorage.c \
	  $(SDK_ROOT)/components/libraries/fstorage/nrf_fstorage_sd.c \
	  $(SDK_ROOT)/components/libraries/experimental_memobj/nrf_memobj.c \
	  $(SDK_ROOT)/components/libraries/pwr_mgmt/nrf_pwr_mgmt.c \
	  $(SDK_ROOT)/components/libraries/experimental_section_vars/nrf_section_iter.c \
	  $(SDK_ROOT)/components/libraries/strerror/nrf_strerror.c \
	  $(SDK_ROOT)/components/libraries/uart/retarget.c \
	  $(SDK_ROOT)/components/libraries/util/sdk_mapped_flags.c \
	  $(SDK_ROOT)/components/boards/boards.c \
	  $(SDK_ROOT)/components/drivers_nrf/clock/nrf_drv_clock.c \
	  $(SDK_ROOT)/components/drivers_nrf/common/nrf_drv_common.c \
	  $(SDK_ROOT)/components/drivers_nrf/gpiote/nrf_drv_gpiote.c \
	  $(SDK_ROOT)/components/drivers_nrf/uart/nrf_drv_uart.c \
	  $(SDK_ROOT)/components/drivers_nrf/timer/nrf_drv_timer.c \
	  $(SDK_ROOT)/external/segger_rtt/SEGGER_RTT.c \
	  $(SDK_ROOT)/external/segger_rtt/SEGGER_RTT_Syscalls_GCC.c \
	  $(SDK_ROOT)/external/segger_rtt/SEGGER_RTT_printf.c \
	  $(SDK_ROOT)/components/ble/common/ble_advdata.c \
	  $(SDK_ROOT)/components/ble/ble_advertising/ble_advertising.c \
	  $(SDK_ROOT)/components/ble/common/ble_conn_params.c \
	  $(SDK_ROOT)/components/ble/common/ble_conn_state.c \
	  $(SDK_ROOT)/components/ble/common/ble_srv_common.c \
	  $(SDK_ROOT)/components/ble/nrf_ble_gatt/nrf_ble_gatt.c \
	  $(SDK_ROOT)/components/ble/ble_services/ble_nus/ble_nus.c \
	  $(SDK_ROOT)/components/softdevice/common/nrf_sdh.c \
	  $(SDK_ROOT)/components/softdevice/common/nrf_sdh_ble.c \
	  $(SDK_ROOT)/components/softdevice/common/nrf_sdh_soc.c \
	  ./Sources/device/system_nrf52.c \
      ./Sources/aiu_net.c \
      ./Sources/main.c \

	  
	  # Libraries common to all targets
#LIB_FILES += d:/Projects_EclipseNRF/meLinkYou/Sources/Libraries/msgMan/Wolfssl/lib/wolfSSL-Lib.a 

##  C++ source files
#CPPFILES = $(filter %.cpp, $(SRC))
##  C source files
#CFILES = $(filter %.c, $(SRC))
##  Assembly source files
#ASMFILES = $(filter %.s, $(SRC))
#
## Object files
#CPPOBJ = $(CPPFILES:.cpp=.o)
#COBJ = $(CFILES:.c=.o)
#SOBJ = $(ASMFILES:.s=.o)
#OBJ  = $(CPPOBJ) $(COBJ) $(SOBJ)


#-Wno-unused-parameter

COMMONFLAGS = -mcpu=cortex-m4 -mthumb -mabi=aapcs -mfloat-abi=hard -mfpu=fpv4-sp-d16

# Optimization flags
OPT = -Os -g3
#OPT = -O0 -g3 -fthread-jumps -falign-functions -falign-jumps -falign-loops -falign-labels -fcaller-saves -fcrossjumping -funswitch-loops \
#-fcse-follow-jumps -fcse-skip-blocks -fdelete-null-pointer-checks -fdevirtualize -fdevirtualize-speculatively -fno-expensive-optimizations \
#-finline-functions -fno-reorder-blocks -fno-reorder-blocks-and-partition -fno-reorder-functions -Wno-missing-braces \
#-fpredictive-commoning -fgcse-after-reload -ftree-loop-vectorize -ftree-loop-distribute-patterns -fsplit-paths -ftree-slp-vectorize -fvect-cost-model -ftree-partial-pre -fpeel-loops -fipa-cp-clone \
#-fno-strict-aliasing

# Uncomment the line below to enable link time optimization
#OPT += -flto
#OPT = -O1 -g3 -ffunction-sections -fdata-sections -fno-strict-aliasing -fno-builtin -fshort-enums -Wno-unused-parameter 

# C flags common to all targets
CFLAGS += -std=gnu11
CFLAGS += $(OPT)
CFLAGS += $(COMMONFLAGS)
CFLAGS += -DBOARD_PCA10040
CFLAGS += -DBSP_DEFINES_ONLY
#CFLAGS += -DCONFIG_GPIO_AS_PINRESET
CFLAGS += -DFLOAT_ABI_HARD
CFLAGS += -DNRF52 
CFLAGS += -DNRF52832_XXAA
CFLAGS += -DNRF52832
CFLAGS += -DNRF52_PAN_74 -DNRF52_PAN_12 -DNRF52_PAN_15 -DNRF52_PAN_20 -DNRF52_PAN_31
CFLAGS += -DSWI_DISABLE0 -DnRF52_DK -DDEBUG_NRF 


CFLAGS += -DBLE_DFU_APP_SUPPORT
CFLAGS += -DBLE_STACK_SUPPORT_REQD
CFLAGS += -DNRF_SD_BLE_API_VERSION=3
CFLAGS += -DSOFTDEVICE_PRESENT
CFLAGS += -DS132
CFLAGS += -DDEBUG
CFLAGS += -DDEBUG_CONFIG
CFLAGS += -DBOARD_CUSTOM
CFLAGS += -DNRF_LOG_USES_RTT=1
CFLAGS += -Wimplicit-fallthrough

CFLAGS += -Wall
# keep every function in a separate section, this allows linker to discard unused ones
CFLAGS += -ffunction-sections -fdata-sections -fno-strict-aliasing
CFLAGS += -fno-builtin -fshort-enums 

#CFLAGS += -funsigned-char -funsigned-bitfields
#CFLAGS += -fpack-struct -fno-exceptions -Wextra -Wno-expansion-to-defined -Winline -Wpointer-arith -Wredundant-decls -Wshadow


# Compile thumb for cortex-m4 with debug info
#CPPFLAGS  = -g -mthumb -DFREERTOS -DBLE_STACK_SUPPORT_REQD -DBOARD_CUSTOM -DCONFIG_GPIO_AS_PINRESET -DNRF52 -DNRF52832_XXAA -DNRF52_PAN_12 -DNRF52_PAN_15 -DNRF52_PAN_20 -DNRF52_PAN_31 -DNRF52_PAN_36 -DNRF52_PAN_51 -DNRF52_PAN_54 -DNRF52_PAN_55 -DNRF52_PAN_58 -DNRF52_PAN_64 -DNRF52_PAN_74 -DNRF_SD_BLE_API_VERSION=5 -DS132 -DSOFTDEVICE_PRESENT -DSWI_DISABLE0 -DnRF52_DK -DDEBUG_NRF -DDEBUG -DBOARD_PCA10040 -mcpu=cortex-m4 -mfloat-abi=hard -mfpu=fpv4-sp-d16 -Og -fpack-struct -fdata-sections -ffunction-sections -fno-exceptions -fno-rtti -std=c++17
CXXFLAGS  += -O0 -g -mthumb -DNRF52832 -DBLE_STACK_SUPPORT_REQD -DBOARD_CUSTOM
#CXXFLAGS  += -DCONFIG_GPIO_AS_PINRESET
CXXFLAGS  += -DNRF52 -DNRF52832_XXAA -DNRF52_PAN_12 -DNRF52_PAN_15 -DNRF52_PAN_20 -DNRF52_PAN_31 -DNRF52_PAN_36 -DNRF52_PAN_51
CXXFLAGS  += -DNRF52_PAN_54 -DNRF52_PAN_55 -DNRF52_PAN_58 -DNRF52_PAN_64 -DNRF52_PAN_74 -DNRF_SD_BLE_API_VERSION=5 
CXXFLAGS  += -DS132 -DSOFTDEVICE_PRESENT -DSWI_DISABLE0 -DnRF52_DK -DDEBUG_NRF -DDEBUG -DBOARD_PCA10040 -mcpu=cortex-m4 -mfloat-abi=hard -mfpu=fpv4-sp-d16 -Og -fpack-struct -fdata-sections -ffunction-sections -fno-exceptions -fno-rtti -std=c++17

# Assembler flags common to all targets
ASMFLAGS += $(COMMONFLAGS)
ASMFLAGS += -x assembler-with-cpp
ASMFLAGS += -g3
ASMFLAGS += -DBOARD_PCA10040
ASMFLAGS += -DBSP_DEFINES_ONLY
ASMFLAGS += -DCONFIG_GPIO_AS_PINRESET
ASMFLAGS += -DFLOAT_ABI_HARD
ASMFLAGS += -DNRF52
ASMFLAGS += -DNRF52832
ASMFLAGS += -DNRF52832_XXAA
ASMFLAGS += -DNRF52_PAN_74

ASMFLAGS += -DBLE_DFU_APP_SUPPORT
ASMFLAGS += -DBLE_STACK_SUPPORT_REQD
ASMFLAGS += -DNRF_SD_BLE_API_VERSION=3
ASMFLAGS += -DSOFTDEVICE_PRESENT
ASMFLAGS += -DS132
ASMFLAGS += -DDEBUG
ASMFLAGS += -DNRF_LOG_USES_RTT=1
ASMFLAGS += -DSVC_INTERFACE_CALL_AS_NORMAL_FUNCTION
ASMFLAGS += -DSWI_DISABLE0
ASMFLAGS += -D__HEAP_SIZE=0
ASMFLAGS += -D__STACK_SIZE=8192
ASFLAGS = -g -mthumb -mcpu=cortex-m4

# Linker flags
LDFLAGS += -std=gnu11
LDFLAGS += -DNRF_SD_BLE_API_VERSION=3
LDFLAGS += $(OPT)
LDFLAGS += $(COMMONFLAGS)
LDFLAGS += -L $(TEMPLATE_PATH) -T$(LINKER_SCRIPT)
# let linker dump unused sections
LDFLAGS += -Wl,--gc-sections
# use newlib in nano version
LDFLAGS += --specs=nano.specs 


# Add standard libraries at the very end of the linker input, after all objects
# that may need symbols provided by these libraries.
LIB_FILES += -lc -lnosys -lm



#all: $(SRC) $(PROJECT).elf $(PROJECT).hex $(PROJECT).bin
#
#$(PROJECT).bin: $(PROJECT).elf
#	@$(CP) -O binary $(PROJECT).elf $@
#
#$(PROJECT).hex: $(PROJECT).elf
#	@$(CP) -O ihex $(PROJECT).elf $@
#
#$(PROJECT).elf: $(OBJ)
#	@echo Linking
#	@$(LD) $(LDFLAGS) $(OBJ) -o $@
#	@$(OS) -t $(PROJECT).elf
#
#
#$(CPPOBJ): %.o: %.cpp
#	@echo $<
#	@$(CPP) -c $(INC) $(CPPFLAGS) $< -o $@
#
#$(COBJ): %.o: %.c
#	@echo $<
#	@$(CC) -c $(INC) $(CFLAGS) $< -o $@
#
#$(SOBJ): %.o: %.s
#	@echo $<
#	@$(AS) -c $(ASFLAGS) $< -o $@
#
#clean:
#	@rm -f $(PROJECT).elf $(PROJECT).bin $(PROJECT).map $(PROJECT).hex $(PROJECT).lss $(OBJ)
#
#install: $(PROJECT).bin
#	JLink.exe -device nRF52832_xxAA -CommanderScript ./flash.jlink


.PHONY: default help

# Default target - first one defined
default: nrf52832_xxaa

# Print all targets that can be built
help:
	@echo following targets are available:
	@echo		nrf52832_xxaa
	@echo		sdk_config - starting external tool for editing sdk_config.h
	@echo		flash      - flashing binary


TEMPLATE_PATH := $(SDK_ROOT)/components/toolchain/gcc


include $(TEMPLATE_PATH)/Makefile.common

$(foreach target, $(TARGETS), $(call define_target, $(target)))

.PHONY: flash erase

# Flash the program
flash: $(OUTPUT_DIRECTORY)/nrf52832_xxaa.hex
	@echo Flashing: $<
	$(JLINK_DIR)/$(JLINK) -device nrf52832_xxaa f SWD -speed 4000 -autoconnect 1 -CommanderScript cmdFile_flashFW.jlink

# Flash softdevice
flash_softdevice:
	@echo Flashing: s132_nrf52_5.1.0_softdevice.hex
	$(JLINK_DIR)/$(JLINK) -device nrf52832_xxaa f SWD -speed 4000 -autoconnect 1 -CommanderScript cmdFile_flashSD.jlink

erase:
	nrfjprog -f nrf52 --eraseall

SDK_CONFIG_FILE := ./Sources/Config/sdk_config.h
CMSIS_CONFIG_TOOL := $(SDK_ROOT)/external_tools/cmsisconfig/CMSIS_Configuration_Wizard.jar
sdk_config:
	java -jar $(CMSIS_CONFIG_TOOL) $(SDK_CONFIG_FILE)
