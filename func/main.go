class msg{
	public :
	  int time;
	  string text;
}

int func main()  {
	map<string,int>mp;
    while(true){
		msg m1 ;
		m1 = getstatus();

	//	set<string>dic;
		if(!mp[m1.text]){
			//set.insert(m1.text);
			cout<<m1.time<<" "<<m1.text;
			mp[m1.text]=m1.time;
		}else{
			//check prv time 
			int prevtime = mp[m1.text];
			if(abs(prevtime-m1.time)>10){
				mp[m1.text]=m1.time;
                cout<<m1.time<<" "<<m1.text<<endl;
			}
		}
	}

	return 0;
}