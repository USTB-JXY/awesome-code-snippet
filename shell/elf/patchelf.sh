root@xxx:/opt/ssh/openssh-8.2p1# patchelf
syntax: patchelf
  [--set-interpreter FILENAME]
  [--page-size SIZE]
  [--print-interpreter]
  [--print-soname]		Prints 'DT_SONAME' entry of .dynamic section. Raises an error if DT_SONAME doesn't exist
  [--set-soname SONAME]		Sets 'DT_SONAME' entry to SONAME.
  [--set-rpath RPATH]
  [--remove-rpath]
  [--shrink-rpath]
  [--allowed-rpath-prefixes PREFIXES]		With '--shrink-rpath', reject rpath entries not starting with the allowed prefix
  [--print-rpath]
  [--force-rpath]
  [--add-needed LIBRARY]
  [--remove-needed LIBRARY]
  [--replace-needed LIBRARY NEW_LIBRARY]
  [--print-needed]
  [--no-default-lib]
  [--debug]
  [--version]
  FILENAME
root@xxx:/opt/ssh/openssh-8.2p1# patchelf --print-needed scp
libdl.so.2
libutil.so.1
libz.so.1
libcrypt.so.1
libresolv.so.2
libc.so.6
ld-linux-aarch64.so.1

root@xxx:/opt/ssh/openssh-8.2p1# patchelf --print-rpath scp

root@xxx:/opt/ssh/openssh-8.2p1# patchelf --set-rpath '/usr/local/lib:/lib:/usr/local/lib' scp
root@xxx:/opt/ssh/openssh-8.2p1# patchelf --print-rpath scp
/usr/local/lib:/lib:/usr/local/lib
