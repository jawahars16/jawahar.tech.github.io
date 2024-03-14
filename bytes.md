---
layout: page
title: Bytes
permalink: /bytes/
---

These are byte sized posts with a simple illustration.
<div class="bytes">
  {% assign sorted = site.bytes | reverse %}
  {% for post in sorted %}
    <h3>{{post.title}} <a id="{{post.slug}}" class="copy-link" title="Copy link to heading">🏷️</a></h3>
    {% if post.content contains "<!-- more -->" %}
      {{ post.content | split:"<!-- more -->" | first % }}
      <div class="see-more">
        <a href="{{ post.url }}">Read more </a>
      </div>
    {% else %}
    <div>
      {{ post.content }}
    </div>
    {% endif %}
  <img src="{{post.image}}" />
  <hr />
  {% endfor %}
</div>