<div class="container">
  <div class="row">
    <div class="col s12" action="/create" method="POST">
      <h2 class="header">Potemkin never gives up</h2>
      <h4 class="">This secret was deleted, you can never access this information again</h4>
      <input type="hidden" value="{{.salt}}" name="salt" class="_salt"/>
      <div class="row">
        <div class="input-field col s12">
          <textarea class="materialize-textarea _secret_show">{{.secret}}</textarea>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <input id="password" type="password" class="_decrypt_pass validate" placeholder="Password (Optional End-to-End)">
          <label for="password" data-error="Sorry, couldn't decrypt, probably unsupported by your browser."></label>
          <button style="float: right;" type="button" class="waves-effect waves-light btn _decryptor">Decrypt</a>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <a href="/" class="waves-effect waves-light btn">Create another</a>
        </div>
      </div>
    </div>
  </div>
</div>
