---
layout: page
title:  "The birth of non blocking IO."
date:   2023-07-26 14:46:39 +0200
image: "/assets/bytes/non-blocking-io.png"
---
Thread per Request is the classic model to process requests. [Of course some servers use thread pool to reduce thread creation time]. 

But the execution escapes the thread most of the times. Because most of the work (moving the data) doesn't involve CPU. So, we can process more requests while waiting. So the idea of Non-Blocking IO born.