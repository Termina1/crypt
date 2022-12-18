<div class="pure-u-1">
  <form class="pure-form pure-form-stacked _submit_new">
    <legend><b>This link will be deleted after first access</b></legend>
    <input type="text" class="pure-u-1" value="{{.}}">
    <a href="/" class="another pure-button pure-button-primary">Create another</a>
    <div class="container pure-u-1 center">
      <img alt="secret qr code" src="/qr.png?uid={{.}}" />
    </div>
  </form>
</div>
