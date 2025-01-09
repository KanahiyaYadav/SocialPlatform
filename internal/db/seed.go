package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"social/internal/store"
)

var usernames = []string{
	"tech_guru", "foodie_diaries", "wanderlust_jane", "fit_and_fab", "bookworm_emily",
	"artsy_andy", "game_master", "travel_tales", "eco_warrior", "fashionista_kate",
	"music_lover", "movie_buff", "science_geek", "coding_ninja", "pet_lover",
	"yoga_enthusiast", "finance_wizard", "history_buff", "meme_king", "fitness_freak",
	"nature_photographer", "sports_fanatic", "minimalist_mary", "urban_farmer",
	"tech_trends", "startuplife", "influencer_dave", "design_thinker", "social_sam",
	"creative_coder", "adventure_alex", "skincare_expert", "home_chef", "news_junkie",
	"crypto_crusher", "content_creator", "digital_nomad", "community_builder", "artist_june",
}

var titles = []string{
	"How to Start Your Day Right", "Top 10 Travel Destinations for 2024",
	"Mastering the Art of Minimalism", "The Beginner's Guide to Investing",
	"Healthy Recipes for Busy Professionals", "10 Must-Read Books This Year",
	"Travel on a Budget: Tips and Tricks", "The Future of AI in Daily Life",
	"Fitness Hacks You Need to Know", "Sustainable Living: A Practical Guide",
	"Best Photography Tips for Beginners", "The Evolution of Social Media",
	"Gaming Trends to Watch in 2024", "The Art of Meditation",
	"Building a Successful Startup", "The Ultimate Skincare Routine",
	"How to Stay Productive While Working Remotely", "The Magic of Daily Journaling",
	"Exploring Local Food Culture", "How to Create Stunning Visual Art",
}

var contents = []string{
	"Start your mornings with these simple habits to boost your energy and productivity.",
	"Explore these incredible travel destinations that you must visit in 2024!",
	"Learn how to simplify your life and focus on what truly matters.",
	"Discover how to start investing, even with a small budget, and grow your wealth.",
	"Quick and healthy meal ideas for busy professionals on the go.",
	"Expand your horizons with these fascinating books.",
	"Find out how to travel the world without breaking the bank.",
	"How artificial intelligence is changing the way we live and work.",
	"Simple exercises and tips to stay fit and healthy in a busy world.",
	"Practical steps to reduce waste and live a sustainable lifestyle.",
	"Capture breathtaking moments with these beginner-friendly photography tips.",
	"A look at the journey of social media and what lies ahead.",
	"What's next in the gaming world? Here are the trends to watch.",
	"Calm your mind and improve focus with these meditation practices.",
	"Step-by-step guide to turning your startup idea into reality.",
	"Achieve glowing skin with these easy-to-follow skincare tips.",
	"Stay productive and focused while working remotely with these strategies.",
	"Why keeping a journal can transform your mindset and creativity.",
	"Experience the flavors and traditions of local food cultures.",
	"Unleash your creativity with these techniques for making stunning art.",
}

var tags = []string{
	"lifestyle", "health", "fitness", "travel", "photography", "technology", "books",
	"finance", "sustainability", "art", "minimalism", "meditation", "gaming", "startups",
	"skincare", "recipes", "inspiration", "productivity", "socialmedia", "culture",
}

var comments = []string{
	"Great tips! Thanks for sharing.",
	"I totally agree with this.",
	"Can't wait to try this out!",
	"This was really helpful, thanks!",
	"Such an inspiring post!",
	"Do you have more content like this?",
	"Loved this! Keep it up.",
	"Very informative, thank you.",
	"Bookmarking this for later.",
	"This is exactly what I needed to hear today.",
	"Great writing as always.",
	"Thanks for breaking it down so well.",
	"Super useful tips, thanks!",
	"This post really resonated with me.",
	"I learned so much from this!",
	"Do you have any other recommendations?",
	"Amazing content, as usual.",
	"This gave me a new perspective.",
	"Can’t wait to implement this!",
	"You’ve inspired me to start something new.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		err := store.Users.Create(ctx, tx, user)
		if err != nil {
			_ = tx.Rollback()
			log.Println("failed to seed users", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		err := store.Posts.Create(ctx, post)
		if err != nil {
			log.Println("failed to seed posts", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		err := store.Comments.Create(ctx, comment)
		if err != nil {
			log.Println("failed to seed comments", err)
			return
		}
	}
	log.Println("seeded successfully")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			// Password: "password", // commented after the transation is done
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	coms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		coms[i] = &store.Comment{
			UserID:  users[rand.Intn(len(users))].ID,
			PostID:  posts[rand.Intn(len(posts))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return coms
}
