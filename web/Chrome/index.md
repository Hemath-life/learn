1. unsubscript the youtube channels in youtubewebsite
   ```
  (async function iife() {
  const UNSUBSCRIBE_DELAY_TIME = 1100; // Increased delay time
  const channels = Array.from(document.querySelectorAll('ytd-channel-renderer'));

  console.log(`${channels.length} channels found.`);
  let ctr = 0;

  for (const channel of channels) {
    const unsubscribeButton = channel.querySelector('.yt-spec-button-shape-next');
    if (unsubscribeButton) {
        console.log(unsubscribeButton)
      unsubscribeButton.click();
      await runAfterDelay(() => {
        const dialog = document.querySelector('.ytd-dialog');
        if (dialog) {
          dialog.querySelector('.close-button').click();
        }
        ctr++;
        if (ctr === channels.length) {
          console.log('All channels unsubscribed!');
        }
      }, UNSUBSCRIBE_DELAY_TIME);
    }
  }
})();

function runAfterDelay(fn, delay) {
  return new Promise((resolve) => {
    setTimeout(() => {
      fn();
      resolve();
    }, delay);
  });
}

```


## hide some sections in youtbe websties
```js
(async function iife() {
  // Get the comments section element
  const commentsSection = document.querySelector('.ytd-comments');

  // If the comments section exists, hide it
  if (commentsSection) {
    commentsSection.style.display = 'none';
  }
})();
```
