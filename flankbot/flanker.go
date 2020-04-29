package flankbot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rodzy/flanker-discordbot/config"
)

//FlankerID is the id for FlankBot
var FlankerID string
var flankSession *discordgo.Session

//FlankStart func to init the connection
func FlankStart() {
	flankSession, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Defintion of the bot discord user
	us, err := flankSession.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	//Setting the user id to the bot
	FlankerID = us.ID

	flankSession.AddHandler(StateHandler)
	flankSession.AddHandler(MessageHandler)

	err = flankSession.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Flanker is running!!!")

	//Setting a signal to stop the bot (CTRL+C)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	flankSession.Close()

}

//MessageHandler to handle all the commands and text from the pdf
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotCom) {
		if m.Author.ID == FlankerID {
			return
		}
	}
	/*-------HELP--------*/
	if m.Content == "$help" {
		//Message introduction
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Git is the open source distributed version control system that facilitates GitHub activities on your laptop or desktop.\n To start just write: ``$<Command>``",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Install",
					Value:  "```html\ncommand -> <$Install>```",
					Inline: false,
				},
				{
					Name:   "Create repositories",
					Value:  "```html\ncommand -> <$Create>```",
					Inline: false,
				},
				{
					Name:   "Configure tooling",
					Value:  "```html\ncommand -> <$Config>```",
					Inline: false,
				},
				{
					Name:   "The .gitignore file",
					Value:  "```html\ncommand -> <$Ignore>```",
					Inline: false,
				},
				{
					Name:   "Branches",
					Value:  "```html\ncommand -> <$Branches>```",
					Inline: false,
				},
				{
					Name:   "Make changes",
					Value:  "```html\ncommand -> <$Changes>```",
					Inline: false,
				},
				{
					Name:   "Synchronize changes",
					Value:  "```html\ncommand -> <$Sync>```",
					Inline: false,
				},
				{
					Name:   "Redo commits",
					Value:  "```html\ncommand -> <$Redo>```",
					Inline: false,
				},
				{
					Name:   "Git Information",
					Value:  "```html\ncommand -> <$Info>```",
					Inline: false,
				},
				{
					Name:   "Tutorial",
					Value:  "```html\ncommand -> <$Tuto>```",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Want to learn more about using GitHub and Git? Email the Training Team or visit our web site for learning event schedules and private class availability.\nservices@github.com\nhttps://services.github.com/",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Flankerbot - Git Commands",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*-------INSTALL--------*/
	if m.Content == "$Install" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Get git on your computer",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "GitHub for Windows",
					Value:  "https://windows.github.com/",
					Inline: false,
				},
				{
					Name:   "GitHub for Mac",
					Value:  "https://mac.github.com/",
					Inline: false,
				},
				{
					Name:   "Git for All Platforms",
					Value:  "https://git-scm.com/",
					Inline: false,
				},
				{
					Name:   "*",
					Value:  "Git distributions for Linux and POSIX systems are available on the official Git SCM web site.",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - Install",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*-------BRANCHES--------*/
	if m.Content == "$Branches" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Branches are an important part of working with Git. Any commits you make will be made on the branch you're currently “checked out” to. Use ``git status`` to see which branch that is.",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Creates a new branch",
					Value:  "```shell\n$ git branch [branch-name]```",
					Inline: false,
				},
				{
					Name:   "Switches to the specified branch and updates the working directory",
					Value:  "```shell\n$ git checkout [branch-name]```",
					Inline: false,
				},
				{
					Name:   "Combines the specified branch’s history into the current branch. This is usually done in pull requests, but is an important Git operation.",
					Value:  "```shell\n$ git merge [branch]```",
					Inline: false,
				},
				{
					Name:   "Deletes the specified branch",
					Value:  "```shell\n$ git branch -d [branch-name]```",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - Branches",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*------CREATE--------*/
	if m.Content == "$Create" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "When starting out with a new repository, you only need to do it once; either locally, then push to GitHub, or by cloning an existing repository.",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Turn an existing directory into a git repository",
					Value:  "```shell\n$ git init```",
					Inline: false,
				},
				{
					Name:   "Clone (download) a repository that already exists on GitHub, including all of the files, branches, and commits",
					Value:  "```shell\n$ git clone [url]```",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - Create repositories",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*------Config-------*/
	if m.Content == "$Config" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Configure user information for all local repositories",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Sets the name you want attached to your commit transactions",
					Value:  "```shell\n$ git config --global user.name [name]```",
					Inline: false,
				},
				{
					Name:   "Sets the email you want attached to your commit transactions",
					Value:  "```shell\n$ git config --global user.email [email address]```",
					Inline: false,
				},
				{
					Name:   "Enables helpful colorization of command line output",
					Value:  "```shell\n$ git config --global color.ui auto```",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - Configure tooling",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*------Git ignore-------*/
	if m.Content == "$Ignore" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Sometimes it may be a good idea to exclude files from being tracked with Git. This is typically done in a special file named ``.gitignore`` . You can find helpful templates for ``.gitignore`` files at github.com/github/gitignore.",
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - The .gitgnore file",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*-------CHANGES--------*/
	if m.Content == "$Changes" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Browse and inspect the evolution of project files",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Lists version history for the current branch",
					Value:  "```shell\n$ git log```",
					Inline: false,
				},
				{
					Name:   "Lists version history for a file, including renames",
					Value:  "```shell\n$ git log --follow [file]```",
					Inline: false,
				},
				{
					Name:   "Shows content differences between two branches",
					Value:  "```shell\n$ git diff [first-branch]...[second-branch]```",
					Inline: false,
				},
				{
					Name:   "Outputs metadata and content changes of the specified commit",
					Value:  "```shell\n$ git show [commit]```",
					Inline: false,
				},
				{
					Name:   "Snapshots the file in preparation for versioning",
					Value:  "```shell\n$ git add [file]```",
					Inline: false,
				},
				{
					Name:   "Records file snapshots permanently in version history",
					Value:  "```shell\n$ git commit -m [descriptive message]```",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - Make changes",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*------Redo-------*/
	if m.Content == "$Redo" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Erase mistakes and craft replacement history",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Undoes all commits after [commit], preserving changes locally",
					Value:  "```shell\n$ git reset [commit]```",
					Inline: false,
				},
				{
					Name:   "Discards all history and changes back to the specified commit",
					Value:  "```shell\n$ git reset --hard [commit]```",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - Redo commits",
			Footer: &discordgo.MessageEmbedFooter{
				Text: "CAUTION! Changing history can have nasty side effects. If you need to change commits that exist on GitHub (the remote), proceed with caution. If you need help, reach out at github.community or contact support.",
			},
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*------Info-------*/
	if m.Content == "$Info" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Some intresting facts",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Sets the name you want attached to your commit transactions",
					Value:  "**git:** an open source, distributed version-control system\n\n **GitHub:** a platform for hosting and collaborating on Git repositories\n\n **commit:** a Git object, a snapshot of your entire repository compressed into a SHA\n\n **branch:** a lightweight movable pointer to a commit\n\n **clone:** a local version of a repository, including all commits and branches\n\n **remote:** a common repository on GitHub that all team member use to exchange their changes\n\n **fork:** a copy of a repository on GitHub owned by a different user\n\n **pull request:** a place to compare and discuss the differences introduced on a branch with reviews, comments, integrated tests, and more\n\n **HEAD:** representing your current working directory, the HEAD pointer can be moved to different branches, tags, or commits when using ``git checkout``",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://arccwiki.uwyo.edu/images/1/19/GitHub_Flow_steps.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - Information",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*------Sync-------*/
	if m.Content == "$Sync" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Synchronize your local repository with the remote repository on GitHub.com",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Downloads all history from the remote tracking branches",
					Value:  "```shell\n$ git fetch```",
					Inline: false,
				},
				{
					Name:   "Combines remote tracking branch into current local branch",
					Value:  "```shell\n$ git merge```",
					Inline: false,
				},
				{
					Name:   "Uploads all local branch commits to GitHub",
					Value:  "```shell\n$ git push```",
					Inline: false,
				},
				{
					Name:   "Updates your current local working branch with all new commits from the corresponding remote branch on GitHub. ``git pull`` is a combination of ``git fetch`` and ``git merge``",
					Value:  "```shell\n$ git pull```",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "Git Cheat Sheet - Synchronize changes",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	/*------Tutorial-------*/
	if m.Content == "$Tuto" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Just in case - Quick setup\n **REMEMBER: YOU NEED TO CHANGE EVERYTHING THAT'S ON SQUARE BRACKETS**",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Create a new repository",
					Value:  "```shell\n$ echo # [repo-name] >> README.md\n$ git init git add README.md\n$ git commit -m [first commit]\n$ git remote add origin https://github.com/[user]/[repo-name].git\n$ git push -u origin master```",
					Inline: false,
				},
				{
					Name:   "Pushing an existing repository from the command line",
					Value:  "```shell\n$ git remote add origin https://github.com/[user]/[repo-name].git\n$ git push -u origin master```",
					Inline: false,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     "GitHub Repo - Quick setup",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

}

//StateHandler for the state of the bot
func StateHandler(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateListeningStatus("$help")
}
