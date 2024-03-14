# GoSplit + Timecode Formatter
---
Quick and easy tool, build it using the go compiler and it should just work.
Timestamp files should be a CSV formatted as follows [Start Time, End Time] the formats for the start and end time should be 00:00:00 style. Thats Hours::Minutes::Seconds.

All files are named output### the first output000 file is a blank file with no length so you can delete it or ignore it whatever you choose. This tool assumes that it is in the same directory as FFMPEG as well as the timestamp file and source MP3.
