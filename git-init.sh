git config user.email coderdba@coder.com
git config user.name coderdba

echo "# docker" >> README.md
git init
#git add README.md
git add .
git commit -m "first commit"
git remote add origin https://github.com/coderdba-coding-org/docker.git
git push -u origin master
