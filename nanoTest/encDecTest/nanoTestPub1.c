#include <stdio.h>
#include <string.h>
#include <stdint.h>
#include <unistd.h>
#include <nanomsg/nn.h>
#include <nanomsg/pubsub.h>

typedef struct msgBuf_s {
    uint16_t msgType;
    char buf[128];
} msgBuf_t;

typedef struct linkInfo_s {
    uint8_t port;
    uint8_t linkState;
} linkInfo_t;

int main()
{
    int i, rv, sock, sndBufSize;
    linkInfo_t linkInfo;
    msgBuf_t msgBuf;
    char sndBuf[128];
    
    sock = nn_socket(AF_SP, NN_PUB);
    if (sock < 0) {
        printf("Failed to create pub socket\n");
        return -1;
    }
    rv = nn_bind(sock, "ipc:///tmp/test.ipc");
    if (rv < 0) {
        printf("Failed to bind pub socket\n");
        return -1;
    }
    sndBufSize = 1024*1024;
    rv = nn_setsockopt(sock, NN_PUB, NN_SNDBUF, &sndBufSize, sizeof(int));
    sleep(1);
    linkInfo.port = 2;
    linkInfo.linkState = 0;
    msgBuf.msgType = 1;
    memcpy(msgBuf.buf, &linkInfo, sizeof(linkInfo_t));
    memcpy(sndBuf, &msgBuf, sizeof(msgBuf_t));
    for (i = 0; i < 1; i++) {
        rv = nn_send(sock, sndBuf, sizeof(linkInfo_t)+sizeof(uint16_t), 0);
        if (rv < sizeof(linkInfo_t)+sizeof(uint16_t)) {
            printf("Failed to send entire msg\n");
            return -1;
        }
    }
    return 0;
}
