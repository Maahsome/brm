# Binary Release Manager (brm)

With so many new tools being released as binary artifacts, some of the other packagers seem like a little bit of overkill.  This will simply go to the source (github, gitlab) releases and download the binary artifact and symlink it into `/usr/local/bin`.

Also be able to switch between binary versions quickly through a symlink change.  This can be useful when you want your `ctl` controller application to match the server api version you are working on.
