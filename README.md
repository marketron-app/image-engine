<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/marketron-app/app">
    <img src="media/marketron-cropped.png" alt="Logo" height="50">
  </a>
<h3 align="center">Marketron Image Engine</h3>

  <p align="center">
        Marketron Image Engine is a service, used in <a href="https://marketron.app">Marketron application</a>. It's main purpose is to generate screenshot of the given URL, transform it to fit into template, and paste it into given template.
    <br />
    
  </p>
</div>


## About The Project
This service is built with the intention to be as basic as possible. This means that this service provides an HTTP endpoint, which other services can call.

## Setup
The best way to set up the project is by using Docker. Run the following command to start up the service:
```bash
 docker run -it --rm \
 -e AWS_ACCESS_KEY_ID=<INSERT_YOUR_AWS_ACCESS_KEY> \
 -e AWS_SECRET_ACCESS_KEY=<INSERT_YOUR_AWS_SECRET_KEY> \
 -e AWS_REGION=<INSERT_AWS_REGION> \
 -e AWS_S3_BUCKET=<INSERT_BUCKET_NAME> \
 -p "3000:3000" \
 tavsec/marketron-image-engine:v1.0.0
```

### Environmental variables
| Variable               | Required | Description                                                                        |
|------------------------|----------|------------------------------------------------------------------------------------|
| AWS_ACCESS_KEY_ID      | yes      | AWS access key                                                                     |
| AWS_SECRET_ACCESS_KEY  | yes      | AWS secret key                                                                     |
| AWS_REGION             | yes      | AWS region                                                                         |
| AWS_S3_BUCKET          | yes      | AWS S3 Bucket, on which the generated images will be uploaded                      |
| AWS_ENDPOINT           | no       | Can be set to use other S3-compatible storages.                                    |
| METRIC_HEADERS_ENABLED | no       | Setting this variable to "true" will enable metric headers. Default is disabled.   |
| CRAWLER_TIMEOUT        | no       | How long (in seconds) the crawler will wait for the website. Default is 10 seconds |

## Usage
Application consists of one main endpoint, and one healthcheck endpoint.

### Generate image
To generate new image, use `GET /image` request, like so:
```bash
curl --location -g --request GET \
   'http://localhost:3000/image? \
    url=https://www.marketron.app& \
    templateImage=https://marketron-exports-images.s3.eu-central-1.amazonaws.com/a6f937aa-b53d-4a8f-9c0d-b70e8413fc7e.png& \
    coordinates[0][x]=88& \
    coordinates[0][y]=126& \
    coordinates[1][x]=125& \
    coordinates[1][y]=466& \
    coordinates[2][x]=650& \
    coordinates[2][y]=448& \
    coordinates[3][x]=597& \
    coordinates[3][y]=138& \
    viewportWidth=1920& \
    viewportHeight=1080' \
--header 'Content-Type: multipart/form-data'
```

#### Parameters
**`url`**: URL of the website, for which you want to generate the mockup.  
**`templateImage`**: Cutout image, on which the screenshot of the website will be pasted on. Must have transparent cutout, in PNG format.  
**`coordinates`**: Array of coordinates (in pixels), where the cutout is on the templateImage. Must be in the following order: top-left, bottom-left, bottom-right, top-right.    
**`viewportWidth`**: Width of the screenshot.  
**`viewportHeight`**: Height of the screenshot.

#### Response
Endpoint will return filename of the generated image, which is uploaded to the S3 bucket.

### Healthcheck
You can use `GET /` or `GET /health` to use as a healthcheck.

### Metrics headers
If you set the `METRIC_HEADERS_ENABLED` to `true`, the `/image` endpoint will pass the following response headers:

| Header                         | Description                                                        |
|--------------------------------|--------------------------------------------------------------------|
| X-Marketron-Metric-Crawler     | How long (in ms) crawler needed to get the website screenshot.     |
| X-Marketron-Metric-Transformer | How long (in ms) application needed to warp screenshot to template |
| X-Marketron-Metric-Uploader    | How long (in ms) application needed to upload image                |


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/marketron-app/image-engine.svg?style=for-the-badge
[contributors-url]: https://github.com/marketron-app/image-engine/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/marketron-app/image-engine.svg?style=for-the-badge
[forks-url]: https://github.com/marketron-app/image-engine/network/members
[stars-shield]: https://img.shields.io/github/stars/marketron-app/image-engine.svg?style=for-the-badge
[stars-url]: https://github.com/marketron-app/image-engine/stargazers
[issues-shield]: https://img.shields.io/github/issues/marketron-app/image-engine.svg?style=for-the-badge
[issues-url]: https://github.com/marketron-app/image-engine/issues
[license-shield]: https://img.shields.io/github/license/marketron-app/image-engine.svg?style=for-the-badge
[license-url]: https://github.com/marketron-app/image-engine/blob/main/LICENSE
