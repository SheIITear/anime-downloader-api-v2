# anime-downloader-api-v2
Search, download and compress anime on fly to small size without losing much from quality.

Installation and running:
>move shelltear.loli to golang src, or set gopath to the directory.

>run 'go get -u -v -f all' to install all other go dependencies.

>install youtube-dl and ffmpeg with needed features.

Usage:
>send get request to localhost:8080/getid/:name, where :name is anime name, this will give you an id you need to provide to the downloader.

>send get request to localhost:8080/getlinks/:id/:onlyinfo, where :id is the id from earlier step and :onlyinfo is a boolean, to either download or just print the infomation.

TODO:
   
   -add json output
   
   -cleaner code

## Disclaimer
When downloading anime, users are subject to country-specific software distribution laws. This is not designed to enable illegal activity. We do not promote piracy nor do we allow it under any circumstances. You should own an original copy of every content downloaded through this tool. Please take the time to review copyright and video distribution laws and/or policies for your country before proceeding.
