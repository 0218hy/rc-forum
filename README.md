# Residential College (RC) Forum
Residential College (RC) web forum is a centralized web-based platform to support communication, coordination and community interaction within RC. It merges announcements, issue reporting and day-to-day communication (Buy / Sell /Give and Open Jio) into a single web forum.

---

## Motivation
As communication is fragmented across Telegram, email and word-of-mouth, this results in loss of information, difficulty tracking issues and poor searchability.

---

## AI usage
AI was used solely for mock data generation.

---

## Tech Stack

**Backend**
- Go (Chi)
- PostgreSQL
- Goose (database migrations)
- SQLC (type-safe query generation)
- JWT authentication

**Frontend**
- React
- TypeScript
- Vite

--- 

## Project Overview

The project follows a decoupled architecture:
- Backend and frontend are implemented as separate components.
- Backend APIs are fully functional and testable independently.
- Frontend focuses on UI structure and basic data presentation.
- Full frontend–backend integration is partially completed at the time of submission.

---

## User Manual

### Setup (Backend):
1.	Clone the project repository
2.	Configure environment variable (Database URL and JWT_SECRET)
3.	Ensure PostgreSQL is running
4.	Run database migrations using Gosse: Goose up 
5.	Start backend server: go run cmd/*.go

### Setup (Frontend):
1.	Navigate to the frontend directory
2.	Install dependencies: npm install
3.	Start the development server: npm run dev

--- 

## Limitation
-	Authentication is implemented in the backend but not integrated into the frontend UI
-	Comment creation are UI only 
-	Not all backend features (register, login) are exposed through frontend integrations
-	End-to-end session handling is not demonstrated via frontend

--- 
## Intended integration
These integrations are partially implemented and were planned for further development
-	Frontend communicates with backend REST APIs
-	JWT access tokens are attached to authenticated request
-	Post and comments are dynamically rendered from backend responses
-	Role-based access control governs functionality.

---

## Reflection
Two months ago, I could not imagine being able to complete what is now in front of me. What felt impossible at the start of this assignment has gradually become something tangible. This assignment, which deliberately forced us to learn without relying on AI, pushed me far beyond my comfort zone and fundamentally changed how I approach learning with AI. 
Learning without AI meant turning to a more traditional, old-school approach. Watching multiple Youtube videos, reading documentations and piecing together solutions from different sources. Unlike getting a single, fast, AI-generated answer, I had to actively decide which approaches to adopt and which to discard. Over time, this helped me develop a much deeper understanding of the tools I was using and make a reasoning behind each of my choices. When I encounter errors now, I no longer blindly throw into ChatGPT. Instead, I read the error messages and usually know where the issue is and how to solve it.
The initial learning curve was extremely extremely steep. It took me more than two weeks just to grasp the basic of Go, React and SQL, as well as to set up my development environment in VS Code and configure PostgreSQL. I spent countless hours watching Youtube Videos and reading documentations whenever unfamiliar concepts appeared. However, despite being slower at the beginning, this method of learning proved to be more efficient in the long run. I can now confidently start a new project and begin coding without AI assistance. The decision I made, ranging from data schema design to architectural choices were entirely my own, which made the learning experience feel genuine and personal. 
This process also taught me that getting lost is not a waste of time. Reading “wrong” solutions or exploring dead ends eventually contributed to my understanding and helped me make better decisions later. I genuinely enjoyed the experience of navigating this maze and gradually finding correct solutions. I am grateful for this assignment as it not only taught me technical skills but also reshaped how I view and use AI. Even if I do not get into CVWO, this experience has been invaluable.
One key regret and learning point is that I should have started smaller and built the project incrementally. Although I had planned an MVP, it was still too ambitious for someone who had never built a web application before. In hindsight, I should have begun with a simple data schema and basic login functionality, then add features while developing the frontend and backend together. As an inexperienced developer, I initially believed I could complete the entire backend first, then move on to the frontend and finally connect them. I now understand that backend and frontend development are deeply interdependent. As a result, I was unable to fully showcase all the backend features I implemented through the frontend.
Through this assignment, I also discovered that I enjoy the backend development far more than frontend work. Designing data models, structuring APIs and making architectural decision were both challenging and rewarding. I only wish I had more time to complete the project. Overall, this assignment was tough but deeply enriching. It has significantly strengthened both my technical foundation and my confidence as a learner. 
