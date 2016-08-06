function str2ab(str) {
  var buf = new ArrayBuffer(str.length*2); // 2 bytes for each char
  var bufView = new Uint16Array(buf);
  for (var i=0, strLen=str.length; i<strLen; i++) {
    bufView[i] = str.charCodeAt(i);
  }
  return buf;
}

function ab2str(buf) {
  return String.fromCharCode.apply(null, new Uint16Array(buf));
}

function initalizeEncrypt(form) {
  form.addEventListener("submit", encryptSecret);
}

function initializeDecryptor(decryptor) {
  decryptor.addEventListener("click", decryptSecret);
}

var cipherType = {
  name: "AES-CTR",
  length: 256,
};

function keyFromPass(pass) {
  var bufPass = str2ab(pass);
  return crypto.subtle.importKey("raw", bufPass, { name: "PBKDF2" }, false, ["deriveKey"])
    .then(function(key) {
      return crypto.subtle.deriveKey({
        name: "PBKDF2",
         salt: new Uint8Array(16),
         iterations: 1000,
         hash: {name: "SHA-1"},
      }, key, cipherType, false, ["encrypt", "decrypt"]);
    });
}

function encrypt(secret, key) {
  var bufSecret = str2ab(secret);
  return crypto.subtle.encrypt({
    name: "AES-CTR",
    counter: new Uint8Array(16),
    length: 128
  }, key, bufSecret).then(function(cipher) {
    return ab2str(cipher);
  });
}

function decrypt(cipher, key) {
  var bufCipher = str2ab(cipher);
  return crypto.subtle.decrypt({
    name: "AES-CTR",
    counter: new Uint8Array(16),
    length: 128,
  }, key, bufCipher).then(function(bufPlain) {
    return ab2str(bufPlain);
  });
}

var cipherText = "";

function decryptSecret(ev) {
  var secret = document.querySelector("._secret_show");
  var pass = document.querySelector("._decrypt_pass");
  pass.classList.remove("invalid");

  var originalCipher = cipherText || secret.value;
  cipherText = originalCipher;
  keyFromPass(pass.value).then(function(key) {
    return decrypt(originalCipher, key);
  }).then(function(plaintext) {
    secret.value = plaintext;
  }).catch(function(err) {
    pass.classList.add("invalid");
  });
}

function encryptSecret(ev) {
  var pass = document.querySelector("._encrypt_pass");
  pass.classList.remove("invalid")
  var secret = document.querySelector("._create_secret");
  if (pass.value) {
    keyFromPass(pass.value).then(function(key) {
      return encrypt(secret.value, key);
    }).then(function(cipher) {
      secret.value = cipher;
      ev.target.removeEventListener('submit', encryptSecret);
      ev.target.submit();
    }).catch(function(error) {
      pass.classList.add("invalid")
      pass.value = "";
    });
    ev.preventDefault();
  }
}


(function() {
  document.addEventListener("DOMContentLoaded", function() {
    var form = document.querySelector('._submit_new');
    var decryptor = document.querySelector('._decryptor');

    if (form) {
      initalizeEncrypt(form);
    }

    if (decryptor) {
      initializeDecryptor(decryptor);
    }
  });
})();
