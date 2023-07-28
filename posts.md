---
layout: page
title: Blog
permalink: /blog/
---

<div class="custom-post">
  <ul>
  {% for post in site.posts %}
    <li>
    <div class="title">
      <div class="heading">
        {%- assign date_format = site.minima.date_format | default: "%b %-d, %Y" -%}
        <span class="post-meta">{{ post.date | date: date_format }}</span>
        <h3>
          <a class="post-link" href="{{ post.url | relative_url }}">
            {{ post.title | escape }}
          </a>
        </h3>
      </div>
      </div>
      {{ post.excerpt | strip_html | strip_newlines | truncate: 360 }}
      <a href="{{ post.url | relative_url }}">Read More</a>
    </li>
  {% endfor %}
  </ul>
</div>