<script>
  function onSubmit(token) {
    document.getElementById("show-form").submit();
  }
</script>
<script src='https://www.google.com/recaptcha/api.js' async></script>
<div class="container">
  <form action="/show" class="pure-form" id="show-form" method="GET">
    <input type="hidden" name="uid" value="{{.uid}}" />
    <button class="g-recaptcha pure-button pure-button-primary"
        data-sitekey="{{.clientKey}}"
        data-callback='onSubmit'
        data-action='submit'>Show secret</button>
  </form>
</div>
