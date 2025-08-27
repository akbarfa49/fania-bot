FROM postgres:latest

# Install necessary packages
RUN apt-get update \
    && apt-get install -y build-essential cron git make gcc postgresql-server-dev-all  \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install pg_cron extension
RUN git clone https://github.com/citusdata/pg_cron.git \
    && cd pg_cron \
    && make \
    && make install
