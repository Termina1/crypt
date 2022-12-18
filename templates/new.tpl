<div class="pure-u-1">
  <form class="pure-form pure-form-stacked _submit_new" action="/create" method="POST">
    <fieldset>
      <legend>We will generate a link for one-time sharing</legend>
      <input type="hidden" value="" name="salt" class="_salt"/>
      <div class="pure-u-1">
        <textarea class="pure-input-1 _create_secret _content_area secret" name="secret" rows="10"
          placeholder="Enter the secret you want to share..."></textarea>
      </div>
      <div class="pure-u-1-2">
        <label for="password-e2e">Password (Optional End-to-End)</label>
        <input type="password" id="password-e2e" class="pure-u-1 _encrypt_pass" placeholder="Enter password" />
        <span class="error" hidden>Sorry, couldn't encrypt, probably unsupported by your browser.</span>
      </div>
    </fieldset>
    <button type="submit" class="pure-button pure-button-primary">Submit</button>
  </form>
</div>
<script>init();</script>
