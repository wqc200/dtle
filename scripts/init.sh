#! /usr/bin/env bash

# chkconfig: 2345 99 01
# description: Dtle daemon

### BEGIN INIT INFO
# Provides:          dtle
# Required-Start:    $all
# Required-Stop:     $remote_fs $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Start dtle at boot time
### END INIT INFO

# this init script supports three different variations:
#  1. New lsb that define start-stop-daemon
#  2. Old lsb that don't have start-stop-daemon but define, log, pidofproc and killproc
#  3. Centos installations without lsb-core installed
#
# In the third case we have to define our own functions which are very dumb
# and expect the args to be positioned correctly.

DTLE_OPTS=

USER=dtle
GROUP=dtle

if [ -r /lib/lsb/init-functions ]; then
    source /lib/lsb/init-functions
fi

if [ -z "$STDOUT" ]; then
    STDOUT=/dev/null
fi
if [ ! -f "$STDOUT" ]; then
    mkdir -p `dirname $STDOUT`
fi

if [ -z "$STDERR" ]; then
    STDERR=/var/log/dtle/dtle.log
fi
if [ ! -f "$STDERR" ]; then
    mkdir -p `dirname $STDERR`
fi

OPEN_FILE_LIMIT=65536

function pidofproc() {
    if [ $# -ne 3 ]; then
        echo "Expected three arguments, e.g. $0 -p pidfile daemon-name"
    fi

    if [ ! -f "$2" ]; then
        return 1
    fi

    local pidfile=`cat $2`

    if [ "x$pidfile" == "x" ]; then
        return 1
    fi

    if ps --pid "$pidfile" | grep -q $(basename $3); then
        return 0
    fi

    return 1
}

function killproc() {
    if [ $# -ne 3 ]; then
        echo "Expected three arguments, e.g. $0 -p pidfile signal"
    fi

    pid=`cat $2`

    kill -s $3 $pid
}

function log_failure_msg() {
    echo "$@" "[ FAILED ]"
}

function log_success_msg() {
    echo "$@" "[ OK ]"
}

# Process name ( For display )
name=dtle

# Daemon name, where is the actual executable
daemon=INSTALL_PREFIX_MAGIC/usr/bin/dtle

# pid file for the daemon
pidfile=/var/run/dtle/dtle.pid
piddir=`dirname $pidfile`

if [ ! -d "$piddir" ]; then
    mkdir -p $piddir
    chown $USER:$GROUP $piddir
fi

# Configuration file
config=INSTALL_PREFIX_MAGIC/etc/dtle/dtle.conf

# If the daemon is not there, then exit.
[ -x $daemon ] || exit 5

case $1 in
    start)
        # Checked the PID file exists and check the actual status of process
        if [ -e $pidfile ]; then
            pidofproc -p $pidfile $daemon > /dev/null 2>&1 && status="0" || status="$?"
            # If the status is SUCCESS then don't need to start again.
            if [ "x$status" = "x0" ]; then
                log_failure_msg "$name process is running"
                exit 0 # Exit
            fi
        fi

        # Bump the file limits, before launching the daemon. These will carry over to
        # launched processes.
        ulimit -n $OPEN_FILE_LIMIT
        if [ $? -ne 0 ]; then
            log_failure_msg "set open file limit to $OPEN_FILE_LIMIT"
        fi

        log_success_msg "Starting the process" "$name"
        if which start-stop-daemon > /dev/null 2>&1; then
            start-stop-daemon --chuid $USER:$GROUP --start --quiet --pid-file $pidfile --exec $daemon -- -pid-file $pidfile -config $config $DTLE_OPTS >>$STDOUT 2>>$STDERR &
        else
            su -s /bin/sh -c "nohup $daemon server -pid-file $pidfile -config $config $DTLE_OPTS >>$STDOUT 2>>$STDERR &" $USER
        fi
        log_success_msg "$name process was started"
        ;;

    stop)
        # Stop the daemon.
        if [ -e $pidfile ]; then
            pidofproc -p $pidfile $daemon > /dev/null 2>&1 && status="0" || status="$?"
            if [ "$status" = 0 ]; then
                if killproc -p $pidfile SIGTERM && /bin/rm -rf $pidfile; then
                    log_success_msg "$name process was stopped"
                else
                    log_failure_msg "$name failed to stop service"
                fi
            fi
        else
            log_failure_msg "$name process is not running"
        fi
        ;;

    reload)
        # Reload the daemon.
        if [ -e $pidfile ]; then
            pidofproc -p $pidfile $daemon > /dev/null 2>&1 && status="0" || status="$?"
            if [ "$status" = 0 ]; then
                if killproc -p $pidfile SIGHUP; then
                    log_success_msg "$name process was reloaded"
                else
                    log_failure_msg "$name failed to reload service"
                fi
            fi
        else
            log_failure_msg "$name process is not running"
        fi
        ;;

    restart)
        # Restart the daemon.
        $0 stop && sleep 3 && $0 start
        ;;

    status)
        # Check the status of the process.
        if [ -e $pidfile ]; then
            if pidofproc -p $pidfile $daemon > /dev/null; then
                log_success_msg "$name Process is running"
                exit 0
            else
                log_failure_msg "$name Process is not running"
                exit 1
            fi
        else
            log_failure_msg "$name Process is not running"
            exit 3
        fi
        ;;

    version)
        $daemon version
        ;;

    *)
        # For invalid arguments, print the usage message.
        echo "Usage: $0 {start|stop|restart|status|version}"
        exit 2
        ;;
esac
