<div class="container">
  <div class="row">
    <form class="col s12 _submit_new" action="/create" method="POST">
      <h2 class="header">Potemkin never gives up</h2>
      <h5 class="">We will generate a link for one-time sharing</h5>
      <input type="hidden" value="" name="salt" class="_salt"/>
      <div class="row">
        <div class="input-field col s12">
          <textarea class="materialize-textarea _create_secret _content_area" name="secret" placeholder="Enter the secret you want to share..."></textarea>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <input id="password" type="password" class="_encrypt_pass validate" placeholder="Password (Optional End-to-End)">
          <label for="password" data-error="Sorry, couldn't encrypt, probably unsupported by your browser."></label>
        </div>
      </div>
      <div class="row">
        <div class="input-field col s12">
          <button type="submit" class="waves-effect waves-light btn">Submit</button>
        </div>
      </div>
    </form>
  </div>
</div>
