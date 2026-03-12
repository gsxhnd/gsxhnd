#include <iostream>
#include <netinet/in.h>
#include <sys/socket.h>
#include <thread>
#include <vector>

void client_thread(int thread_id) {
  // 1. 创建 socket
  int sockfd = socket(AF_INET, SOCK_STREAM, 0);
  // 2. 设置服务器地址
  sockaddr_in serv_addr{};
  // ... (设置地址族、端口、IP)
  serv_addr.sin_family = AF_INET;
  serv_addr.sin_addr.s_addr = ntohl(INADDR_LOOPBACK);
  serv_addr.sin_port = ntohs(8080);

  std::cout << "socketfd:" << sockfd << std::endl;

  // 3. 连接服务器
  if (connect(sockfd, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) == -1) {
    std::cerr << "Thread " << thread_id << " 连接失败" << std::endl;
    return;
  }
  std::cout << "Thread " << thread_id << " 连接成功" << std::endl;

  // 4. 发送和接收数据（根据测试需求编写）
  std::string message = "Hello from thread " + std::to_string(thread_id);
  std::cout << "Thread send msg " << message << std::endl;
  if (send(sockfd, message.c_str(), message.size(), 0) < 0) {
    perror("send failed:");
  };
  // std::cout << "Thread send msg status" << a  << std::endl;

  char buffer[1024];
  ssize_t n = recv(sockfd, buffer, sizeof(buffer), 0);
  if (n > 0) {
    std::cout << "Thread " << thread_id
              << " 收到回复: " << std::string(buffer, n) << std::endl;
  }

  // 5. 关闭连接
  close(sockfd);
}

int main() {
  const int num_threads = 100; // 模拟100个并发连接
  std::vector<std::thread> threads;

  for (int i = 0; i < num_threads; ++i) {
    std::cerr << "Thread " << i + 1 << " starting to connect to server"
              << std::endl;
    threads.emplace_back(client_thread, i + 1);
    std::this_thread::sleep_for(std::chrono::milliseconds(0));
  }

  for (auto &t : threads) {
    // 等待所有线程结束
    t.join();
  }
  return 0;
}