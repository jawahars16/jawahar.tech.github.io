---
layout: page
title: Videos
permalink: /videos/
---
The channel focus on introducing software concepts - that cover area like Languages and Frameworks, Infrastructure, DevOps, OOPS and many more. Technical breadth is essential for full stack development and it is important criteria in most full stack dev interviews. 

This channel helps you to become the jack of all concepts. 😑

It is difficult to understand the details of all technical concepts in shorter period of time. And it also tiresome to watch hourly long videos to understand the concepts. The objective of this channel is to introduce the technical concepts and explain the crux of those concepts within 5 to 10 minutes. It is up to you to go ahead and understand those concepts in details, if it interest you. The idea is to keep the video as simple and short as possible to just introduce the concepts and make the viewers understand the purpose and use cases for such concepts. 

[Subscribe to the YouTube hannel.](https://www.youtube.com/@jawahar.selvaraj?sub_confirmation=1){:target="_blank"}
<div class="video-container">
{% for video in site.videos %}
<div class="video"> 
  <iframe src="{{video.link}}" width="350" height="195" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
</div>
{% endfor %}
</div>

[Watch more videos on YouTube.](https://www.youtube.com/@jawahar.nutshell/videos){:target="_blank"}