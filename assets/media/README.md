add an mp4 file here for temp usage
run this for create hls slices
```

ffmpeg -i input.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls 1/hls/index.m3u8
```

