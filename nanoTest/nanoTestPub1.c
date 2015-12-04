#include <stdio.h>
#include <unistd.h>
#include <nanomsg/nn.h>
#include <nanomsg/pubsub.h>

int main()
{
    int i, rv, sock, sndBufSize;
    char msg[] = "Test message - 1";
    
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
    for (i = 0; i < 10000; i++) {
        rv = nn_send(sock, msg, 17, 0);
        if (rv < 17) {
            printf("Failed to send entire msg\n");
            return -1;
        }
    }
    return 0;
}
