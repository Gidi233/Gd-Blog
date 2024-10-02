# Copyright Gidi233 <qpbtyfh@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/Gidi233/Gd-Blog.

FROM BASE_IMAGE

WORKDIR /opt/Gd-Blog

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/Gd-Blog

COPY Gd-Blog /opt/Gd-Blog/bin/

ENTRYPOINT ["/opt/Gd-Blog/bin/Gd-Blog"]
CMD ["-c", "/etc/Gd-Blog/Gd-Blog.yaml"]
