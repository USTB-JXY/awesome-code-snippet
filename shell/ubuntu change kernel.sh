root@xxx:~# uname -r
5.4.0-156-generic
root@xxx:~# vim /etc/default/grub
root@xxx:~# 
root@xxx:~# grep submenu /boot/grub/grub.cfg
submenu 'Advanced options for Ubuntu' $menuentry_id_option 'gnulinux-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
root@xxx:~# grep gnulinux /boot/grub/grub.cfg
   set default="gnulinux-5.4.0-153-generic-advanced-663dcb85-232a-4102-be78-b8a671828dbf"
menuentry 'Ubuntu' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-simple-663dcb85-232a-4102-be78-b8a671828dbf' {
submenu 'Advanced options for Ubuntu' $menuentry_id_option 'gnulinux-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-156-generic' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-156-generic-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-156-generic (recovery mode)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-156-generic-recovery-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-153-generic' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-153-generic-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-153-generic (recovery mode)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-153-generic-recovery-663dcb85-232a-4102-be78-b8a671828dbf' {
root@xxx:~# grep menuentry /boot/grub/grub.cfg
if [ x"${feature_menuentry_id}" = xy ]; then
  menuentry_id_option="--id"
  menuentry_id_option=""
export menuentry_id_option
menuentry 'Ubuntu' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-simple-663dcb85-232a-4102-be78-b8a671828dbf' {
submenu 'Advanced options for Ubuntu' $menuentry_id_option 'gnulinux-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-156-generic' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-156-generic-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-156-generic (recovery mode)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-156-generic-recovery-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-153-generic' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-153-generic-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-153-generic (recovery mode)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-153-generic-recovery-663dcb85-232a-4102-be78-b8a671828dbf' {
root@xxx:~# grep menuentry /boot/grub/grub.cfg
if [ x"${feature_menuentry_id}" = xy ]; then
  menuentry_id_option="--id"
  menuentry_id_option=""
export menuentry_id_option
menuentry 'Ubuntu' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-simple-663dcb85-232a-4102-be78-b8a671828dbf' {
submenu 'Advanced options for Ubuntu' $menuentry_id_option 'gnulinux-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-156-generic' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-156-generic-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-156-generic (recovery mode)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-156-generic-recovery-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-153-generic' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-153-generic-advanced-663dcb85-232a-4102-be78-b8a671828dbf' {
	menuentry 'Ubuntu, with Linux 5.4.0-153-generic (recovery mode)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-5.4.0-153-generic-recovery-663dcb85-232a-4102-be78-b8a671828dbf' {
root@xxx:~# 



root@xxx:~# grep 'menuentry \|submenu ' /boot/grub/grub.cfg | cut -f2 -d "'"
Ubuntu
Advanced options for Ubuntu
Ubuntu, with Linux 5.4.0-156-generic
Ubuntu, with Linux 5.4.0-156-generic (recovery mode)
Ubuntu, with Linux 5.4.0-153-generic
Ubuntu, with Linux 5.4.0-153-generic (recovery mode)
root@xxx:~# vim /etc/default/grub
root@xxx:~# update-grub
Sourcing file `/etc/default/grub'
Sourcing file `/etc/default/grub.d/init-select.cfg'
Generating grub configuration file ...
Found linux image: /boot/vmlinuz-5.4.0-156-generic
Found initrd image: /boot/initrd.img-5.4.0-156-generic
Found linux image: /boot/vmlinuz-5.4.0-153-generic
Found initrd image: /boot/initrd.img-5.4.0-153-generic
Warning: Please don't use old title `Ubuntu, with Linux 5.4.0-153-generic' for GRUB_DEFAULT, use `Advanced options for Ubuntu>Ubuntu, with Linux 5.4.0-153-generic' (for versions before 2.00) or `gnulinux-advanced-663dcb85-232a-4102-be78-b8a671828dbf>gnulinux-5.4.0-153-generic-advanced-663dcb85-232a-4102-be78-b8a671828dbf' (for 2.00 or later)
done
root@xxx:~# vim /etc/default/grub
root@xxx:~# update-grub
Sourcing file `/etc/default/grub'
Sourcing file `/etc/default/grub.d/init-select.cfg'
Generating grub configuration file ...
Found linux image: /boot/vmlinuz-5.4.0-156-generic
Found initrd image: /boot/initrd.img-5.4.0-156-generic
Found linux image: /boot/vmlinuz-5.4.0-153-generic
Found initrd image: /boot/initrd.img-5.4.0-153-generic
done

Connecting to 192.168.1.57:22...
Connection established.
To escape to local shell, press 'Ctrl+Alt+]'.

Welcome to Ubuntu 20.04.5 LTS (GNU/Linux 5.4.0-153-generic x86_64)


  System information as of Tue 29 Aug 2023 10:17:41 AM CST

  System load:            1.54
  Usage of /:             41.7% of 98.95GB
  Memory usage:           9%
  Swap usage:             0%
  Processes:              264
  Users logged in:        0
  IPv4 address for ens33: 192.168.1.57
  IPv6 address for ens33: fd66:f61a:ba98:4892:250:56ff:fe93:576b

Last login: Tue Aug 29 10:01:15 2023 from 192.168.2.155
root@xxx:~# 
root@xxx:~# 
root@xxx:~# 
root@xxx:~# uname -r
5.4.0-153-generic
root@xxx:~# 
root@xxx:~# 
root@xxx:~# cat /proc/version
Linux version 5.4.0-153-generic (buildd@bos03-amd64-008) (gcc version 9.4.0 (Ubuntu 9.4.0-1ubuntu1~20.04.1)) #170-Ubuntu SMP Fri Jun 16 13:43:31 UTC 2023
