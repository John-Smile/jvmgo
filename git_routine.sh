echo "================git routine start============"
oldDevBranch=$1
newDevBranch=$2
echo old develop branch $oldDevBranch
echo new develop branch $newDevBranch
git checkout debug
echo "================debug branch status============"
git status
git pull

git checkout master
echo "================master branch status========"
git status
git pull
git merge --no-ff debug 
git push origin master

git checkout $oldDevBranch
echo "================old dev branch status========"
git status
git pull
git checkout debug
git merge --no-ff $oldDevBranch
git push origin debug

git checkout -b $newDevBranch
git push origin $newDevBranch

echo "================git routine end==============" 

