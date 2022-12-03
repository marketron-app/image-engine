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

<h3 align="center">Marketron Image Engine</h3>

  <p align="center">
    Marketron Image Engine is a service, used in <a href="https://marketron.app">Marketron application</a>. It's main purpose is to generate screenshot of the given URL, transform it to fit into template, and paste it into given template.
    <br />
    
  </p>
</div>



<!-- ABOUT THE PROJECT -->
## About The Project


This service is built with the intention to be as basic as possible. This means that this service provides an gRPC endpoint, which other services can call. This service WILL not handle any other logic (authentication, users, ...).
<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started
The quickest way to setup the project is using docker-compose.
```sh
docker-compose up -d --build
```




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
[license-url]: https://github.com/marketron-app/image-engine/blob/master/LICENSE.txt
