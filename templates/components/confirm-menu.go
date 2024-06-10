package components

type ConfirmMenuProps struct {
  Header HeaderProps
  Options []struct{
    Text string
  }
}
var ConfirmMenuTemlplate = "templates/components/confirm-menu.html"
