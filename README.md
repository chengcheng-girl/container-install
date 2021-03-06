# container-install

### Usage

```sh
# 下载并安装container-install, container-install是个golang的二进制工具，直接下载拷贝到bin目录即可, release页面也可下载
$ wget -c https://cuisongliu.oss-cn-beijing.aliyuncs.com/container-install/latest/linux_amd64/container-install && \
    chmod +x container-install && mv container-install /usr/bin
```

### [amd64 下载地址]
[oss 下载地址](https://cuisongliu.oss-cn-beijing.aliyuncs.com/container-install/latest/linux_amd64/container-install)
### [arm64 下载地址]
[oss 下载地址](https://cuisongliu.oss-cn-beijing.aliyuncs.com/container-install/latest/linux_arm64/container-install)


### PreInstall
Download latest [container-install](https://github.com/cuisongliu/container-install/releases/download/v1.2/container-install) on release page.

### Install Docker
install container from url:
   location file:
  ```shell script
   container-install install \
          -d \
          --host 172.16.213.131 \
          --host 172.16.213.132 \
          --user root \
          --passwd admin \
          --registry 127.0.0.1 \
          --registry 127.0.0.2 \
          --lib /var/lib/docker \
          --pkg-url  /root/docker-19.0.3.tgz
```
   remote url:
  ```shell script
   container-install install \
          -d \
          --host 172.16.213.131 \
          --host 172.16.213.132 \
          --user root \
          --passwd admin \
          --registry 127.0.0.1 \
          --registry 127.0.0.2 \
          --lib /var/lib/docker \
          --pkg-url  https://download.docker.com/linux/static/stable/x86_64/docker-18.09.4.tgz
```

### Install Containerd
install container from url:
   location file:
  ```shell script
   container-install install \
          --host 172.16.213.131 \
          --host 172.16.213.132 \
          --user root \
          --passwd admin \
          --registry 127.0.0.1 \
          --registry 127.0.0.2 \
          --lib /var/lib/containerd \
          --pkg-url  /root/containerd-1.2.7.tgz
```
   remote url:
  ```shell script
   container-install install  \
          --host 172.16.213.131 \
          --host 172.16.213.132 \
          --user root \
          --passwd admin \
          --registry 127.0.0.1 \
          --registry 127.0.0.2 \
          --lib /var/lib/containerd \
          --pkg-url  https://github.com/containerd/containerd/releases/download/v1.2.7/containerd-1.2.7.linux-amd64.tar.gz
```

### UnInstall Docker
uninstall container:
  ```shell script
   container-install uninstall  \
          -d \
          --host 172.16.213.131 \
          --host 172.16.213.132 \
          --user root \
          --passwd admin \
          --lib /var/lib/docker
```

### UnInstall Containerd
uninstall container:
  ```shell script
   container-install uninstall \
          --host 172.16.213.131 \
          --host 172.16.213.132 \
          --user root \
          --passwd admin \
          --lib /var/lib/containerd
```

### Login type

default use --password is password login
use  --pk=/root/.ssh/id_rsa is private key login

### Print Download Url
print download url for docker:
 ```shell script
  container-install print -d
```

the docker Newest version is v19.03.0.
ex:

```shell script
cuisongliu@cuisongliu-PC:~$ container-install print -d
https://download.docker.com/linux/static/stable/x86_64/docker-17.03.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.03.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.03.2-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.06.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.06.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.06.2-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.09.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.09.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.12.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.12.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.03.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.03.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.06.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.06.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.06.2-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.06.3-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.0.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.1.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.2.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.3.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.4.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.5.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.6.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.7.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.8.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.9.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.0.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.1.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.10.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.11.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.12.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.13.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.14.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.2.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.3.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.4.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.5.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.6.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.7.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.8.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.9.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.0.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.1.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.10.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.11.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.12.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.13.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.14.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.2.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.3.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.4.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.5.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.6.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.7.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.8.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-rootless-extras-19.03.9.tgz
```

the containerd Newest version is v1.2.7
ex:
```shell script
cuisongliu@cuisongliu-PC:~$ container-install print
https://github.com/containerd/containerd/releases/download/v1.4.3/containerd-1.4.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.4.2/containerd-1.4.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.4.1/containerd-1.4.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.4.0/containerd-1.4.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.4.0-rc.1/containerd-1.4.0-rc.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.4.0-rc.0/containerd-1.4.0-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.4.0-beta.2/containerd-1.4.0-beta.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.4.0-beta.1/containerd-1.4.0-beta.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.4.0-beta.0/containerd-1.4.0-beta.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.9/containerd-1.3.9-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.8/containerd-1.3.8-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.7/containerd-1.3.7-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.6/containerd-1.3.6-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.5/containerd-1.3.5-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.4/containerd-1.3.4-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.3/containerd-1.3.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.2/containerd-1.3.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.1/containerd-1.3.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.0/containerd-1.3.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.0-rc.3/containerd-1.3.0-rc.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.0-rc.2/containerd-1.3.0-rc.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.0-rc.1/containerd-1.3.0-rc.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.0-rc.0/containerd-1.3.0-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.0-beta.2/containerd-1.3.0-beta.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.0-beta.1/containerd-1.3.0-beta.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.3.0-beta.0/containerd-1.3.0-beta.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.14/containerd-1.2.14-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.13/containerd-1.2.13-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.12/containerd-1.2.12-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.11/containerd-1.2.11-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.10/containerd-1.2.10-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.9/containerd-1.2.9-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.8/containerd-1.2.8-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.7/containerd-1.2.7-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.6/containerd-1.2.6-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.5/containerd-1.2.5-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.4/containerd-1.2.4-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.3/containerd-1.2.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.2/containerd-1.2.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.1/containerd-1.2.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.1-rc.0/containerd-1.2.1-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.0/containerd-1.2.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.0-rc.2/containerd-1.2.0-rc.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.0-rc.1/containerd-1.2.0-rc.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.0-rc.0/containerd-1.2.0-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.0-beta.2/containerd-1.2.0-beta.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.0-beta.1/containerd-1.2.0-beta.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.2.0-beta.0/containerd-1.2.0-beta.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.8/containerd-1.1.8-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.7/containerd-1.1.7-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.6/containerd-1.1.6-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.5/containerd-1.1.5-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.4/containerd-1.1.4-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.3/containerd-1.1.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.2/containerd-1.1.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.1/containerd-1.1.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.1-rc.2/containerd-1.1.1-rc.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.1-rc.1/containerd-1.1.1-rc.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.1-rc.0/containerd-1.1.1-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.0/containerd-1.1.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.0-rc.2/containerd-1.1.0-rc.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.0-rc.1/containerd-1.1.0-rc.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.1.0-rc.0/containerd-1.1.0-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.3/containerd-1.0.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.3-rc.0/containerd-1.0.3-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.2/containerd-1.0.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.2-rc.1/containerd-1.0.2-rc.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.2-rc.0/containerd-1.0.2-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.1/containerd-1.0.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.1-rc.0/containerd-1.0.1-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0/containerd-1.0.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-rc.0/containerd-1.0.0-rc.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-beta.3/containerd-1.0.0-beta.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-beta.2/containerd-1.0.0-beta.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-beta.1/containerd-1.0.0-beta.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-beta.0/containerd-1.0.0-beta.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-alpha6/containerd-1.0.0-alpha6-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-alpha5/containerd-1.0.0-alpha5-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-alpha4/containerd-1.0.0-alpha4-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-alpha3/containerd-1.0.0-alpha3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-alpha2/containerd-1.0.0-alpha2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-alpha1/containerd-1.0.0-alpha1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v1.0.0-alpha0/containerd-1.0.0-alpha0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.9/containerd-0.2.9-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.8/containerd-0.2.8-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.7/containerd-0.2.7-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.6/containerd-0.2.6-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.5/containerd-0.2.5-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.4/containerd-0.2.4-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.3/containerd-0.2.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.2/containerd-0.2.2-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.1/containerd-0.2.1-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.2.0/containerd-0.2.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/v0.1.0/containerd-0.1.0-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/0.0.5/containerd-0.0.5-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/0.0.4/containerd-0.0.4-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/0.0.3/containerd-0.0.3-linux-amd64.tar.gz
https://github.com/containerd/containerd/releases/download/0.0.2/containerd-0.0.2-linux-amd64.tar.gz
```
