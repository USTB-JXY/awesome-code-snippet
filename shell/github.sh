vim /etc/hosts

140.82.113.3    github.com
185.199.108.153 assets-cdn.github.com
199.232.69.194  github.global.ssl.fastly.net


140.82.113.3    github.com
185.199.108.153 assets-cdn.github.com
151.101.1.194  github.global.ssl.fastly.net

git submodule update --init --recursive


alias proxy_on="export https_proxy=127.0.0.1:7890 && export http_proxy=127.0.0.1:7890"
alias proxy_off="unset http_proxy https_proxy"

git config --global https.proxy https://127.0.0.1:7890



git log
git add chinese-ebook/*
git commit -m 'add folder'
git push -u origin main
