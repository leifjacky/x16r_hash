#include <stdio.h>
#include <string.h>
#include "x16r.h"

int main(){
        unsigned int buf[]={0x30000000, 0x8d7931ff, 0x930f6459, 0xe0b451d3, 0xb1ce13a3, 0x81ac9e3d, 0xa94ba4f9, 0x7ea255c8,
                        0x00000000, 0x5a52c38f, 0xeba3d259, 0xe4e8c9e4, 0x9796ec86, 0xa2b1d5e6, 0x92dd5d43, 0x7256770f,
                        0x11cdc4c1, 0x5c9c8b27, 0x1d07b449, 0xcfdd4e06
        };
        unsigned char output[32];
        memset(output, 0, 8);
        x16r_hash((char *)buf, (char *)output);
        int i = 0;
        for (i = 0; i < 32; i++){
                printf("%02x", output[32 - 1 - i]);
        }
        printf("\n");
        return 0;
}