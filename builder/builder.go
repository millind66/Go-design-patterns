package builder

import "errors"

type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func NewNotificationBuilder() *NotificationBuilder {

	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}
func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.Subtitle = subtitle
}
func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}
func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}
func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}
func (nb *NotificationBuilder) SetNotType(nottype string) {
	nb.NotType = nottype
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) Build() (*Notification, error) {

	if nb.Title == "" {
		return nil, errors.New("Empty Titile found")
	}
	return &Notification{
		title:    nb.Title,
		subtitle: nb.Subtitle,
		message:  nb.Image,
		icon:     nb.Icon,
		image:    nb.Image,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil

}
