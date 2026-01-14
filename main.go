package main

import (
	"context"
	"fmt"
	"log"

	"redis/zset"
)

func main() {
	// Testing the Zset
	ctx := context.Background()
	zset := zset.NewZset(ctx, "test")

	fmt.Println("Adding members to the zset")
	// Add a member to the zset
	err := zset.Add(ctx, 1, "test1")
	if err != nil {
		log.Fatal(err)
	}
	// Add a member to the zset
	err = zset.Add(ctx, 2, "test2")
	if err != nil {
		log.Fatal(err)
	}

	// Add a member to the zset
	err = zset.Add(ctx, 3, "test3")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Getting the score of a member")
	// Get the score of a member
	score, err := zset.Get(ctx, "test2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)

	fmt.Println("Getting all members of the zset")
	// Get all members of the zset
	all, err := zset.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(all)

	fmt.Println("Getting all members greater than 2")
	// Get all members greater than 2
	greaterThan, err := zset.GetGreaterThan(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greaterThan)

	fmt.Println("Getting all members less than 2")
	// Get all members less than 2
	lessThan, err := zset.GetLessThan(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lessThan)

	fmt.Println("Getting all members between 1 and 3")
	// Get all members between 1 and 3
	between, err := zset.GetBetween(ctx, 1, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(between)

	fmt.Println("Removing a member from the zset")
	// Remove a member from the zset
	err = zset.Remove(ctx, "test2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test2 removed")

	fmt.Println("Removing all members from the zset")
	// Remove all members from the zset
	err = zset.RemoveAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("all members removed")

	// duplicate member
	err = zset.Add(ctx, 4, "test3")
	if err != nil {
		log.Fatal(err)
	}

	// get the score of a member
	score, err = zset.Get(ctx, "test3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)

	// add a duplicate member
	err = zset.Add(ctx, 2, "test3")
	if err != nil {
		log.Fatal(err)
	}

	// get the score of a member
	score, err = zset.Get(ctx, "test3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)

	err = zset.Add(ctx, 10, "test3")
	if err != nil {
		log.Fatal(err)
	}

	// get the score of a member
	score, err = zset.Get(ctx, "test3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)

	fmt.Println("Getting all members of the zset")
	// Get all members of the zset
	all, err = zset.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(all)
}
