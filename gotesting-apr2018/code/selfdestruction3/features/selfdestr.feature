 Feature: Self-destruction setup with 2-Factor authentication
   In order to initialize the self destruction sequence 
   As a human officer of a starship
   I need an additional officer to approve the destruction
    
   Scenario Outline: initialize the self-destruction
     Given there are <number> human officers
     When <number> officers enter the destruct sequence
     Then it is <access> to run self-destruction
     
     Examples:
       | number | access    |
       | 1      | forbidden |
       | 2      | allowed   |
