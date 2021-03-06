FROM golang:1.15

# Ignore APT warnings about not having a TTY
ENV DEBIAN_FRONTEND noninteractive

# install build essentials
RUN apt-get update && \
    apt-get install -y wget build-essential pkg-config --no-install-recommends

# Install ImageMagick deps
RUN apt-get -q -y install libjpeg-dev libpng-dev libtiff-dev \
    libgif-dev libx11-dev --no-install-recommends

ENV IMAGEMAGICK_VERSION=7.0.11-2

RUN cd && \
    curl -L -O https://github.com/strukturag/libde265/releases/download/v1.0.3/libde265-1.0.3.tar.gz && \
    tar zxvf libde265-1.0.3.tar.gz && \
    cd libde265-1.0.3 && \
    ./configure --disable-dec265 --disable-sherlock265 && \
    make && make install && \
    cd .. && \
    rm -rf libde265-1.0.3

RUN cd && \
    curl -L -O  https://github.com/strukturag/libheif/releases/download/v1.3.2/libheif-1.3.2.tar.gz && \
    tar zxvf libheif-1.3.2.tar.gz && \
    cd libheif-1.3.2 && \
    ./configure && \
    make && make install && \
    cd .. && \
    rm -rf libheif-1.3.2

RUN cd && \
	wget https://github.com/ImageMagick/ImageMagick/archive/${IMAGEMAGICK_VERSION}.tar.gz && \
	tar xvzf ${IMAGEMAGICK_VERSION}.tar.gz && \
	cd ImageMagick* && \
	./configure \
	    --without-magick-plus-plus \
	    --without-perl \
	    --disable-openmp \
	    --with-gvc=no \
	    --disable-docs && \
	make -j$(nproc) && make install && \
	ldconfig /usr/local/lib && \
	cd .. && \
	rm -rf ImageMagick*

WORKDIR /go/projects/thumbnail-api
COPY . .

RUN cd thumbnail-api && go install
CMD /go/bin/thumbnail-api