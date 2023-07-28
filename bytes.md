---
layout: page
title: Bytes
permalink: /bytes/
---

These are byte sized posts with a simple illustration.
<div class="bytes">
  {% for post in site.bytes %}
    <h3>{{post.title}}</h3>
    {{post.content}}  
  <img src="{{post.image}}" />
  <hr />
  {% endfor %}
</div>