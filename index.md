---
layout: home
---

<div class="home-title">
  <img src="/assets/profile-circle.png" alt="Jawahar Selvaraj" class="profile-home"/>
  <div>
    <h1>Jawahar Selvaraj</h1>
    <p class="typewriter">I CODE FOR LIVING... </p>
    <p></p>
  </div>
  <p>
    <a href="https://www.linkedin.com/in/jawaharselvaraj">Follow me on LinkedIn</a> . 
    <a href="https://www.youtube.com/@jawahar.nutshell">Subscribe to YouTube channel</a>
  </p>
</div>

---
<h2 class="home-h2">
  <a href="/blog">> Blog</a>
</h2>

<div class="custom-post">
  <ul>
  {% for post in site.posts limit:5 %}
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
      <p>
      {{ post.excerpt | strip_html | strip_newlines | truncate: 360 }}
      <a href="{{ post.url | relative_url }}">Read More...</a>
      </p>
    </li>
  {% endfor %}
  </ul>
</div>

<h2 class="home-h2">
  <a href="/bytes">> Bytes</a>
</h2>


<p>These are some byte sized posts with a simple illustration.</p>
<div class="home-bytes">
  {% assign sorted = site.bytes | reverse %}
  {%- assign date_format = site.minima.date_format | default: "%b %-d, %Y" -%}
  {% for post in sorted limit:3 %}
  <div class="byte">
    <time>{{post.date | date: date_format}}</time>
    <h3>{{post.title}}</h3>
    <a href="/bytes/#{{post.slug}}"><img src="{{post.image}}" /></a>
</div>
{% endfor %}
</div>
<p></p>

---

<h2 class="home-h2">
  <a href="/videos">> Videos</a>
</h2>

<p>The channel focus on introducing software concepts - that cover area like Languages and Frameworks, Infrastructure, DevOps, OOPS and many more. Technical breadth is essential for full stack development and it is important criteria in most full stack dev interviews. </p>
<div class="home-videos">
{% for video in site.videos limit:4 %}
<div class="video"> 
  <iframe src="{{video.link}}" width="340" height="200" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
</div>
{% endfor %}
</div>

---