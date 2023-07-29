---
layout: home
---
<img src="/assets/profile-circle.png" alt="Jawahar Selvaraj" class="profile-home"/>
# Jawahar Selvaraj

I am a Software Engineer passionate in Cloud engineering and Golang programming. I have several years of industry experience and proven ability to design, develop, and deploy scalable and reliable applications. I also have expertise in CI/CD, automation, and application release management.

I am an individual contributor who loves to learn and experiment with new technologies. I am also a strong advocate for open source software and I am always looking for ways to share my knowledge with the community.

[Follow me on LinkedIn](https://www.linkedin.com/in/jawaharselvaraj) . [Subscribe to YouTube channel](https://www.youtube.com/@jawahar.nutshell)

---

## Latest blog posts

<ul class="home-posts">
{% for post in site.posts  limit:3 %}
    <li  class="post">
    <img src="{{post.tileImage}}" />
    {%- assign date_format = site.minima.date_format | default: "%b %-d, %Y" -%}
    <div class="heading">
    <span class="post-meta">{{ post.date | date: date_format }}</span>
    <h3>
        <a class="post-link" href="{{ post.url | relative_url }}">
        {{ post.title | escape }}
        </a>
    </h3>
    </div>
    {{ post.excerpt | strip_html | strip_newlines | truncate: 100 }}
    <a href="{{ post.url | relative_url }}">Read More</a>
    </li>
{% endfor %}
</ul>
<p />

[Show all articles](/blog)

---

## Bytes

These are byte sized posts with a simple illustration.

<div class="home-bytes">
  {% for post in site.bytes limit:3 %}
  <div class="byte">
    <h3>{{post.title}}</h3>
    <img src="{{post.image}}" />
</div>
{% endfor %}
</div>
<p />

[Show all bytes](/bytes)

---

## Videos

<div class="home-videos">
{% for video in site.videos limit:3 %}
<div class="video"> 
  <iframe src="{{video.link}}" width="233" height="155" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
</div>
{% endfor %}
</div>

[Watch more videos on YouTube.](https://www.youtube.com/@jawahar.nutshell/videos){:target="_blank"}