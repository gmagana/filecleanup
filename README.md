# filecleanup
A utility written in Go to delete files. Intended to delete the oldest X logs or backup files.

---
# What filecleanup does

A familiar thing to do on a computer system is to create backups from scripts. These backups must then be erased or cleaned up manually, usually by deleting the oldest backups and keeping the mst recent X backups.

*filecleanup* is a cross-platform utility that obtains a list of these backups, sorts them in memory, and then deletes files in order until the last X item remain. In the above-mentioned example of backup files, this has the effect of deleting the oldest files until only the most recent files remain.

This utility can then be used as part of your backup script so that only the most recent automated backups are kept on your disk.

## Configuration Process

File deletion is a dangerous thing. Before you go off deleting files, you should follow a process to make sure you have your commands configured correctly, so they will delete the correct files in unattended mode.

By default, *filecleanup* works in dry run mode, meaning no files are actually deleted. When you are confident in your configuration and are ready to destroy files, use the `-live-run` flag.

### Suggested Workflow

We suggest you follow the following process to configure your *filecleanup* commands:

1. Execute your command using the various flags (see below), including the `-list-all-files` flag so you can see what the full sorted file list looks like.
1. When you have your command line arguments perfected, use the `-live-run` argument to start deleting files.

---

# Usage

`filecleanup [arguments] <file filter>`

***By default this utility runs in dry run (no files are actually deleted) unless you specify the `-live-run` flag. By default, you cannot shoot yourself in the foot.***

## Arguments

| Argument | Comments| 
| --- | --- |
| `-files-to-keep <int>` | (REQUIRED) The number of files you intend to exist after the deletion takes place |
| `-list-all-files` | Show list of all files in order in which they will be processed |
| `-live-run` | If specified, files are deleted. If not specified, a dry run takes place (no files are deleted) |
| `-order-case-insensitive` | Sorts filenames in case insensitive order |
| `-order-reverse` | Sorts filenames in reverse order |
| `<file filter>` | File naming pattern of the files that should be processed. In essence you can use `*` (matches any number of characters) and `?` (matches a single character) in your file masks. Full documentation for how to format this argument is [here](https://hackage.haskell.org/package/Glob-0.9.2/docs/System-FilePath-Glob.html#v:compile). |
