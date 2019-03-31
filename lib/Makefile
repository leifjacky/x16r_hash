CC=gcc
FLAGS=-std=gnu99 -I. -I./sha3
CFLAGS=-fPIC -pedantic -Wall -Wextra -march=native
LDFLAGS=-shared

TARGET=libx16r.so
SOURCES = x16r.c $(shell echo sha3/*.c)
HEADERS = x16r.h $(shell echo sha3/*.h)
#OBJECTS = $(SOURCES:.c=.o)
OBJECTS = \
    x16r.o \
    ./sha3/sph_hefty1.o \
    ./sha3/sph_fugue.o \
    ./sha3/aes_helper.o \
    ./sha3/sph_blake.o \
    ./sha3/sph_bmw.o \
    ./sha3/sph_cubehash.o \
    ./sha3/sph_echo.o \
    ./sha3/sph_groestl.o \
    ./sha3/sph_jh.o \
    ./sha3/sph_keccak.o \
    ./sha3/sph_luffa.o \
    ./sha3/sph_shavite.o \
    ./sha3/sph_simd.o \
    ./sha3/sph_skein.o \
    ./sha3/sph_whirlpool.o \
    ./sha3/sph_shabal.o \
    ./sha3/hamsi.o \
    ./sha3/sha2.o \
    ./sha3/sph_sha2big.o \

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $(LDFLAGS) -o $(TARGET) $(OBJECTS)

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

clean:
	rm *.o sha3/*.o *.so
