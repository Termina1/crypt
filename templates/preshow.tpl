<script src='https://www.google.com/recaptcha/api.js' async></script>
<div class="container">
  <form action="/show" method="GET">
    <input type="hidden" name="uid" value="{{.}}" />
    <div class="recaptcha">
      <div class="g-recaptcha" data-sitekey="6LeUEycTAAAAAPDsPFGqOZ1j8juPRhONLvHrsdrg"></div>
      <button type="submit" class="waves-effect waves-light btn">Show</button>
    </div>
  </form>
</div>
