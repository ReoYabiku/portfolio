window.onload = function() {
  const menu_about = document.getElementById('menu-about');
  if(menu_about != null) {
    console.log("not null");
  }
  const menu_philosophy = document.getElementById('menu-philosophy');
  const menu_biography = document.getElementById('menu-biography');
  const menu_achievement = document.getElementById('menu-achievement');
  const menu_contact = document.getElementById('menu-contact');

  const menu = document.getElementById('menu');
  const shadow = document.getElementById('shadow');
  const hidden = document.getElementById('hidden');

  console.log(menu.style.display);

  menu_about.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
  menu_philosophy.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
  menu_biography.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
  menu_achievement.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
  menu_contact.addEventListener('click', function() {
    menu.style.display = 'none';
    shadow.style.display = 'none';
    hidden.checked = false;
    location.reload();
  })
}
