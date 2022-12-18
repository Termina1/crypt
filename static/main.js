function str2ab(str) {
  var buf = new ArrayBuffer(str.length * 2); // 2 bytes for each char
  var bufView = new Uint16Array(buf);
  for (var i = 0, strLen = str.length; i < strLen; i++) {
    bufView[i] = str.charCodeAt(i);
  }
  return buf;
}

function ab2str(buf) {
  var bufView = new Uint16Array(buf);
  return bufView.toString();
}

function toRealString(buf) {
  return String.fromCharCode.apply(null, new Uint16Array(buf));
}

function init() {
  var text = document.querySelector("._content_area");
  function resize() {
    text.style.height = "auto";
    text.style.height = text.scrollHeight + 5 + "px";
  }
  /* 0-timeout to get the already changed text */
  function delayedResize() {
    window.setTimeout(resize, 0);
  }
  text.addEventListener("input", resize, false);
  text.focus();
  text.select();
  resize();
}

function deab2str(str) {
  var arr = str.split(",").map(function (i) {
    return parseInt(i);
  });
  var buf = new ArrayBuffer(arr.length * 2);
  var bufView = new Uint16Array(buf);
  arr.forEach(function (v, i) {
    bufView[i] = v;
  });
  return buf;
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

function keyFromPass(pass, salt) {
  var bufPass = str2ab(pass);
  return crypto.subtle
    .importKey("raw", bufPass, { name: "PBKDF2" }, false, ["deriveKey"])
    .then(function (key) {
      return crypto.subtle.deriveKey(
        {
          name: "PBKDF2",
          salt: salt,
          iterations: 1000,
          hash: { name: "SHA-1" },
        },
        key,
        cipherType,
        false,
        ["encrypt", "decrypt"]
      );
    });
}

function encrypt(secret, key) {
  var bufSecret = str2ab(secret);
  return crypto.subtle
    .encrypt(
      {
        name: "AES-CTR",
        counter: new Uint8Array(16),
        length: 128,
      },
      key,
      bufSecret
    )
    .then(function (cipher) {
      return ab2str(cipher);
    });
}

function decrypt(cipher, key) {
  var bufCipher = deab2str(cipher);
  return crypto.subtle
    .decrypt(
      {
        name: "AES-CTR",
        counter: new Uint8Array(16),
        length: 128,
      },
      key,
      bufCipher
    )
    .then(function (bufPlain) {
      return toRealString(bufPlain);
    });
}

var cipherText = "";

function decryptSecret(ev) {
  var secret = document.querySelector("._secret_show");
  var pass = document.querySelector("._decrypt_pass");
  var saltEl = document.querySelector("._salt");
  pass.nextElementSibling.setAttribute("hidden", "");

  var originalCipher = cipherText || secret.value;
  cipherText = originalCipher;
  var salt = saltEl.value.split(",").map(function (i) {
    return parseInt(i);
  });
  salt = new Uint8Array(salt);
  keyFromPass(pass.value, salt)
    .then(function (key) {
      return decrypt(originalCipher, key);
    })
    .then(function (plaintext) {
      secret.value = plaintext;
      secret.style.height = "auto";
      secret.style.height = secret.scrollHeight + "px";
    })
    .catch(function (err) {
      pass.nextElementSibling.removeAttribute("hidden");
    });
}

function encryptSecret(ev) {
  var pass = document.querySelector("._encrypt_pass");
  var secret = document.querySelector("._create_secret");
  var saltEl = document.querySelector("._salt");
  var salt = window.crypto.getRandomValues(new Uint8Array(16));
  pass.nextElementSibling.setAttribute("hidden", "");
  saltEl.value = salt.toString();
  if (pass.value) {
    keyFromPass(pass.value, salt)
      .then(function (key) {
        return encrypt(secret.value, key);
      })
      .then(function (cipher) {
        secret.value = cipher;
        ev.target.removeEventListener("submit", encryptSecret);
        ev.target.submit();
      })
      .catch(function (error) {
        pass.nextElementSibling.removeAttribute("hidden");
        pass.value = "";
        saltEl.value = "";
      });
    ev.preventDefault();
  }
}

(function () {
  document.addEventListener("DOMContentLoaded", function () {
    var form = document.querySelector("._submit_new");
    var decryptor = document.querySelector("._decryptor");
    if (form) {
      initalizeEncrypt(form);
    }

    if (decryptor) {
      initializeDecryptor(decryptor);
    }
  });
})();
