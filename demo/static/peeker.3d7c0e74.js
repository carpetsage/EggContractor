(()=>{const e=document.querySelector("form#peeker"),t=document.querySelector("input#contract-id"),o=document.querySelector("input#coop-code"),n=document.querySelector("button#clear-contract-id");e?.addEventListener("submit",(e=>{e.preventDefault();let n=t?.value.trim().toLowerCase(),r=o?.value.trim().toLowerCase();if(n=n?.replace(/\(.*\)/g,"").trim(),!n||!r)return;const c=`/peek/${encodeURIComponent(n)}/${encodeURIComponent(r)}/`;window.open(c,"_blank","noopener")})),n?.addEventListener("click",(e=>{e.preventDefault(),t&&(t.value="")}))})();