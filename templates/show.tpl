<div class="pure-u-1">
  <form class="pure-form pure-form-stacked">
    <legend><b>This secret was deleted, you can never access this information again</b></legend>
    <fieldset>
      <div class="pure-u-1">
        <textarea class="pure-input-1 _secret_show _content_area secret" name="secret" rows="10"
          placeholder="Enter the secret you want to share...">{{.secret}}</textarea>
      </div>
      <input type="hidden" value="{{.salt}}" name="salt" class="_salt"/>
      <div class="pure-u-1-2">
        <label for="password-e2e">Password (Optional End-to-End)</label>
        <input id="password-e2e" type="password" class="pure-u-1 _decrypt_pass" placeholder="Enter password">
        <span class="pure-form-message error" hidden>Sorry, couldn't decrypt, probably unsupported by your browser.</span>
        <button type="button" class="pure-button pure-button-primary _decryptor another">Decrypt</button>
        <a href="/" class="pure-button another">Create another</a>
      </div>
    </fieldset>
  </form>
</div>
<script>init();</script>
