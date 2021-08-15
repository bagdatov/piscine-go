INTERVIEW_NUM=$(cat streets/Buckingham_Place | grep "SEE INTERVIEW #" | tr -dc '0-9\n')
echo "$INTERVIEW_NUM"
cat interviews/interview-"$INTERVIEW_NUM"
SUSPECTS=$(cat ./vehicles | grep -A5 'L337*' | grep -A4 -B1 'Honda' | grep -A3 -B2 'Blue' | grep -B4 'Height: 6')
echo "$SUSPECTS"
grep "$MAIN_SUSPECT" ./memberships/AAA ./memberships/Delta_SkyMiles ./memberships/Museum_of_Bash_History ./memberships/Terminal_City_Library | wc -l