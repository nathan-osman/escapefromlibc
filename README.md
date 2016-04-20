## escapefromlibc

This tool serves no practical purpose under ordinary circumstances. But it becomes a lifesaver when this accidentally happens:

    rm -rf /lib/ x86_64-linux-gnu/libfancy.so

Oops. An extra space caused the command to wipe `/lib`. Restoring the contents of the directory from a backup may be possible but suddenly none of the standard Unix utilities work:

    # mkdir /something
    mkdir: error while loading shared libraries: libc.so.6
    # vi something.txt
    vi: error while loading shared libraries: libc.so.6

How do we fix this mess? If you installed escapefromlibc, you can breathe a sigh of relief. This handly little utility provides enough functionality to get you back on your feet.

    # elc wget 'http://backup_server/bkp.tar.gz'
    # elc tar xf bkp.tar.gz
    # elc cp bkp/libc.so.6 /lib/x86_64-linux-gnu

escapefromlibc accomplishes all of this _without a dependency on libc_. It uses kernel syscalls directly to work its magic and restore your system to working order.

### Building elc

Build elc is as simple as:

    make
