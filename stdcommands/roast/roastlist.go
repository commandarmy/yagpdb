package roast

import (
	"math/rand"
)

func randomRoast() string {
	return roasts[rand.Intn(len(roasts))]
}

//Roasts copied from roast-bot
var roasts = []string{
	"I'd like to roast you, but it looks like your genetics already did.",
	"You look like someone set your face on fire and then put it out with a hammer.",
	"The only thing attracted to you is gravity",
	"You're not good looking enough to be a model, but you're not smart enough to be anything else",
	"If you'd like to know what sexual position produces the ugliest babies, you should ask your mother.",
	"Can you speak a little louder? I can't hear you over the sound of how stupid you are.",
	"Why are you even talking to me? So your self-esteem can match your IQ?",
	"I'm not insulting you, I'm describing you.",
	"If you hide your big nose and shut your big mouth, people will think you are attractive and well-spoken.",
	"You're so ugly, when your mom dropped you off at school, she got a fine for littering.",
	"Some babies were dropped on their heads, but you were clearly thrown at a wall.",
	"They say opposites attract. If that's so, you will meet someone who is good-looking, intelligent, and cultured.",
	"I didn't hear you. I'm busy ignoring an annoying person.",
	"I don't know what your problem is, but I'll bet it's hard to pronounce.",
	"Please excuse me while I transfer you to someone who speaks Fucktard.",
	"It must take a lot of flexibility to fit your foot in your mouth and your head up your ass at the same time.",
	"I don't have the time nor the crayons to explain things to you",
	"I'd love to keep chatting with you, but I'd rather have AIDS",
	"I bet you swim with a t shirt on",
	"You have all the charm and charisma of a burning orphanage",
	"Your face is so oily that I'm surprised America hasn't invaded yet.",
	"If you were any dumber, someone would need to water you twice a week",
	"If you were on fire and I had a cup of my own piss, I'd drink it",
	"Do you still love nature, despite what it did to you?",
	"The thing I dislike most about your face is that I can see it.",
	"If B.S. was music, you'd be an orchestra.",
	"You look like a before picture.",
	"I've heard farts more intelligent than you.",
	"You have a perfect face for radio.",
	"They say that a million monkeys on a million typewriters will eventually produce the collected works of Shakespeare. If that theory is correct, I believe you will one day say something intelligent.",
	"You look like somebody stepped on a goldfish.",
	"I thought the trash got picked up last night, what are you still doing here?",
	"Looking the way you do must save a lot of money on Halloween.",
	"I'd love to continue talking with you but my favorite commercial is on tv",
	"I'd love to keep chatting with you, but right now I have to do literally anything else.",
	"Did you get a bowl of soup with that haircut?",
	"If you don't like what I say about you, it would be a good idea to improve yourself.",
	"Does being that ugly require a license?",
	"You could throw a rock at the ground and miss",
	"There's no one in this world like you. Or at least I hope so",
	"Even a duck is smarter than you.",
	"Did you cancel your barbecue?  Because your grill is messed up",
	"Some people make millions.  You make memes.",
	"Did you forget to wipe or is that your natural scent?",
	"I missed you this week, but my aim is improving.",
	"I'm surprised you've made it this far without being eaten.",
	"Your body looks like your head is inflating a water balloon.",
	"Your mother was a hamster.",
	"How do you make an idiot wait? I'll tell you later",
	"If balls were dynamite, you wouldn't have enough to kill a fish.",
	"I'd like to roast you, but I'm too busy judging your choices.",
	"You are the worst part of everybody's day.",
	"If your face were scrambled it would improve your looks.",
	"I hope you don't feel the way you look.",
	"In the book of Who's Who, you are listed as What's That?",
	"It's surprising to me that a pig's bladder on a stick has gotten so far in life.",
	"Sorry.  I'm on the toilet and I can only deal with one shit at a time.",
	"If you fell into a river it would be unfortunate, but if anyone pulled you out it would be a disaster.",
	"You are the discount version of whatever celebrity you look like.",
	"When you go to the dentist, the dentist needs anesthetic.",
	"The fact that you are still alive is evidence that natural disasters are poorly distributed.",
	"You are so dumb you can't fart and chew gum at the same time.",
	"I was going to give you a nasty look, but I see you already have one.",
	"Me think'st thou are a general offense and every man should beat thee.",
	"I don't try to explain myself to idiots like you.  I'm not the Fucktard Whisperer.",
	"Your mom circulates like a public key, servicing more requests than HTTP.",
	"Your mom is so fat and dumb, the only reason she opened her email is because she heard it contained spam.",
	"Your mom is so fat, she has to iron her pants in the driveway.",
	"Your face invites a slap.",
	"The only way you could get laid is if you crawled up a chicken's ass and waited.",
	"There's somebody out there for everybody. For you, it's a psychiatrist.",
	"It's a parents job to raise their children right. So looking at you, it's no wonder your dad quit after just one day.",
	"Maybe you should try to eat make-up to improve your ugly personality.",
	"You're so ugly your face makes onions cry!",
	"There is something gross on your face, never mind, it is your face.",
	"You're not completely useless, you can always serve as a bad example.",
	"You're so ugly, when you got robbed, the robbers made you wear their masks.",
	"You're the reason they invented double doors!",
	"You stare at frozen juice cans because they say, \"concentrate\".",
	"They say people get what they deserve. In your case it's a participation trophy.",
	"I'd offer you some gum but your smiles got plenty of it.",
	"I fart in your general direction",
	"There aren't enough curse words to use for you",
	"If I had a gun with three bullets and I was in a room with Hitler, bin Laden, Toby, and you, I'd shoot you thrice",
	"You're a basic bitch",
	"You're the AT&T of people",
	"You're the opposite of batman, you're the human equivalent of a tennis elbow, you're a pizza burn on the roof of the world's mouth",
	"If you were spice, you'd be flour",
	"You're a negative ten on a scale of 0 to 5",
	"You look like a pipe cleaner with eyes",
	"Somewhere out there is a tree, tirelessly producing oxygen so you can breathe, I think you owe it an apology",
	"It's okay to be a furry, it's not okay to be you",
	"You've aged",
	"You're an open book written for very dumb children",
	"I always hear \"punch me in the face\" when you're speaking, but it's usually subtext",
	"Don't talk out loud. You lower the IQ of the whole street.",
	"I'm not questioning your honor. I'm denying its existence.",
	"You must have been born on the highway, because that's where most accidents happen.",
	"If you were a trophy at the end of a race, I'd walk backwards.",
	"Anybody who told you to just be yourself couldn't have given you worse advice.",
	"The jerk store called and they're running out of you!",
	"If I had a dollar for every brain you didn't have, I'd have one dollar.",
	"You're so pretentious, you'd eat a worm if it had a French name.",
	"At Cornell University they have an incredible piece of scientific equipment known as the tunnelling electron microscope. Now, this microscope is so powerful that by firing electron you can actually see images of the atom, the infinitesimally minute building blocks of our universe, If I were using that microscope right now, I still wouldn't be able to locate your brain.",
	"Your garden is overgrown and your cucumbers are soft",
	"I have had wet shits better than your personality",
	"Here's a phone, call someone who cares.",
	"You should photoshop your life with better decisions.",
	"You're a grimy little pimp",
	"Fuck you",
	"If genius skips a generation, your children will be brilliant.",
	"Whatever doesn't kill you, disappoints me.",
	"You're my favorite person… besides every other person I've ever met.",
	"I hope your wife brings a date to your funeral.",
	"You're about as useful as a screen door on a submarine",
	"I can't wait to spend the whole life without you.",
	"Take my lowest priority and put yourself beneath it.",
	"I don't hate you, but if you were drowning, I would give you a high five",
	"Unless your name is Google, stop acting like you know everything",
	"I didn't mean to offend you, but it was a huge plus",
	"There is someone out there for everyone. For you, it's a therapist.",
	"Sorry, I can't think of an insult dumb enough for you to understand.",
	"I would smack you, but I'm against animal abuse.",
	"You are the reason there are instructions on a shampoo bottle.",
	"I envy people who have never met you.",
	"Remember when I asked for your opinion? Me neither.",
	"Whoever told you to be yourself gave you bad advice.",
	"You are a pizza burn on the roof of the world's mouth.",
	"I don't know what makes you so stupid, but it works.",
	"You have such a beautiful face, but let's put a bag over that personality.",
	"You're impossible to underestimate.",
	"I don't have the time or crayons to explain this to you.",
	"I told my therapist about you; She didn't believe me.",
	"I could eat alphabet soup and shit out a smarter statement than whatever you just said.",
	"You ass must be pretty jealous of all the shit that comes out of your mouth.",
	"Grab a straw, because you suck.",
	"Have a nice day... somewhere else.",
	"You look like a \"before\" picture.",
	"Yes, I'm fully vaccinated, but I still won't hang out with you.",
	"If ignorance is bliss, you must be the happiest person on earth.",
	"You're more disappointing than an unsalted pretzel.",
	"You're the reason I prefer animals to people.",
	"I'm not insulting you, I'm describing you",
	"Why is it acceptable for you to be an idiot but not for me to point it out?",
	"The last time I saw something like you… I flushed.",
	"You're like a plunger. You like to bring up old shit.", 
	"I treasure the time I don't spend with you.",
	"You're as useless as the \"ueue\" in \"queue\"",
	"You're the reason the gene pool needs a lifeguard.",
	"I don't know what your problem is, but I'm guessing it's hard to pronounce.",
	"Talking to you is like stepping on a leaf in autumn and hearing no crunch- disappointment.",
}
