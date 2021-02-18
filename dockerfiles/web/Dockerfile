FROM library/golang

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y bash \
    zsh \
    vim

ENV APP_DIR $GOPATH/src/app
RUN mkdir -p $APP_DIR

# Godep for vendoring
RUN go get github.com/tools/godep
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/lib/pq
RUN go get -u github.com/astaxie/beego/cache/redis
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gomodule/redigo/redis
RUN go get gopkg.in/mgo.v2
RUN go get -u golang.org/x/crypto/bcrypt

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std


# Set the entrypoint
#ENTRYPOINT (cd $APP_DIR)

# Compile the binary and statically link
#RUN cd $APP_DIR && CGO_ENABLED=0 godep go build -ldflags '-d -w -s'

# zsh
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true
ADD conf/.zshrc /root/.zshrc

##WkHtmlToPdf
#RUN curl -fkL "https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.3/wkhtmltox-0.12.3_linux-generic-amd64.tar.xz" -o wkhtmltopdf.tar.xz && \
#    tar -xJf wkhtmltopdf.tar.xz && \
#    mv wkhtmltox/bin/wkhtmltopdf /usr/bin/wkhtmltopdf && \
#    mv wkhtmltox/bin/wkhtmltoimage /usr/bin/wkhtmltoimage && \
#    chmod +x /usr/bin/wkhtmltopdf
#    mkdir -p /var/www/symfony/bin && \
#    ln -s /usr/bin/wkhtmltopdf /var/www/symfony/bin/wkhtmltopdfs

WORKDIR "/go/src/app"