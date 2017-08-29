FROM debian

MAINTAINER Hamish Ogilvy

RUN apt-get update && apt-get install -y \
 zip \
 poppler-utils \
 wv \
 unrtf \
 tidy \
 lynx
 
EXPOSE 8000

COPY docd /
CMD ["/docd"]
