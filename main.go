package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
)

func main()  {
	fmt.Println("proto buff running")

	person := &Person{
		Name: "Sagar",
		Age: 27,
		SocialFollower: &SocialFollowers{
			Youtube: 32,
			Twitter: 24,
		},
	}

	data, err := proto.Marshal(person)

	if err !=  nil {
		log.Fatal("marshalling error: ", err)
	}

	fmt.Println(data)

	newPerson := &Person{}

	err = proto.Unmarshal(data, newPerson)

	if err != nil {
		log.Fatal("unmarshalling error: ", err)
	}

	fmt.Println(newPerson.GetName())
	fmt.Println(newPerson.GetAge())
	fmt.Println(newPerson.GetSocialFollower().GetYoutube())
	fmt.Println(newPerson.GetSocialFollower().GetTwitter())

}
