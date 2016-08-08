<script src='https://www.google.com/recaptcha/api.js' async></script>
<div class="container">
  <form action="/show" method="GET">
    <input type="hidden" name="uid" value="{{.uid}}" />
    <div class="recaptcha">
      <div class="g-recaptcha" data-sitekey="{{.clientKey}}"></div>
      <button type="submit" class="waves-effect waves-light btn">Show</button>
    </div>
  </form>
</div>
