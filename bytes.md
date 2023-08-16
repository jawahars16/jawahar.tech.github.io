---
layout: page
title: Bytes
permalink: /bytes/
---

These are byte sized posts with a simple illustration.
<div class="bytes">
  {% assign sorted = site.bytes | reverse %}
  {% for post in sorted %}
    <h3>{{post.title}}</h3>
    {{post.content}}  
  <img src="{{post.image}}" />
  <hr />
  {% endfor %}
</div>