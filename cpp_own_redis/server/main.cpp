#include <assert.h>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <netinet/in.h>
#include <ostream>
#include <sys/socket.h>
#include <unistd.h>

const int PORT = 8080;
const int BUFFER_SIZE = 1024;
const size_t k_max_msg = 4096;

static int32_t read_full(int fd, char *buf, size_t n) {
  while (n > 0) {
    ssize_t rv = read(fd, buf, n);
    if (rv <= 0) {
      return -1; // error, or unexpected EOF
    }
    assert((size_t)rv <= n);
    n -= (size_t)rv;
    buf += rv;
  }
  return 0;
}

static int32_t one_request(int connfd) {
  // 4 bytes header
  char rbuf[4 + k_max_msg];
  errno = 0;
  int32_t err = read_full(connfd, rbuf, 4);
  return err;
}

int main() {
  int fd = socket(AF_INET, SOCK_STREAM, 0);
  if (fd < 0) {
    perror("socket created failed");
    return -1;
  }

  int opt = 1;
  if (setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt))) {
    perror("setsockopt failed:");
    close(fd);
    return -1;
  }

  sockaddr_in address{};
  int addrlen = sizeof(address);
  address.sin_family = AF_INET;
  address.sin_addr.s_addr = INADDR_ANY;
  address.sin_port = htons(PORT);

  // if (bind(fd, (struct sockaddr*)&address, addrlen) < 0) {
  if (bind(fd, reinterpret_cast<struct sockaddr *>(&address), addrlen) < 0) {
    perror("bind failed");
    close(fd);
    return -1;
  }

  if (listen(fd, 1) < 0) {
    perror("listen failed:");
    return -1;
  }

  // int new_socket;
  // sockaddr_in client_addr;
  // socklen_t client_addr_len = sizeof(client_addr);
  // if ((new_socket = accept(fd, (struct sockaddr *)&client_addr,
  // &client_addr_len)) < 0) {
  //     perror("accept failed");
  //     close(fd);
  // }
  //
  // if ((new_socket = accept(fd, reinterpret_cast<struct sockaddr *>(&address),
  // &client_addr_len)) < 0) {}

  while (true) {
    struct sockaddr_in client_addr = {};
    socklen_t client_addr_len = sizeof(client_addr);
    int connfd = accept(fd, reinterpret_cast<struct sockaddr *>(&client_addr),
                        &client_addr_len);
    if (connfd < 0) {
      perror("accept failed:");
      continue;
    }
    std::cout << "get client connect" << std::endl;

    while (true) {
      /* code */
      std::cout << "connfd" << connfd << std::endl;
      int32_t err = one_request(connfd);
      if (err) {
        break;
      }
    }

    // char buffer[BUFFER_SIZE] = {0};
    // ssize_t valread = read(connfd, buffer, BUFFER_SIZE);
    // std::cout << buffer << std::endl;

    // const char *resp = "resp from server";
    // send(connfd, resp, strlen(resp), 0);
    close(connfd);
  }

  return 0;
}
